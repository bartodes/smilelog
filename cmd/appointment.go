package cmd

import (
	"fmt"
	"log"
	"time"

	. "github.com/bartodes/smilelog/internals/models"
	. "github.com/bartodes/smilelog/internals/services"
	"github.com/spf13/cobra"
)

var appointment Appointment
var visit Visit

var appointmentCmd = &cobra.Command{
	Use:   "appointment",
	Short: "Manage appointments",
}

var appointmentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an appointment",
	Run: func(cmd *cobra.Command, args []string) {
		if err := PatientExists(appointment.PatientID, db); err != nil {
			log.Fatal(err)
		}

		appointment.ScheduledFor = time.Now().UTC().Format(time.RFC3339)

		overlap, err := CheckAppointmentOverlap(appointment.ScheduledFor, appointment.DurationMinutes, db)

		if overlap {
			log.Fatalf("the appointment for patient '%d' could not be scheduled: %v", appointment.PatientID, ErrAppointmentOverlap)
			// GetAvailableScheduleForAppointment()
		}

		a, err := CreateAppointment(appointment, db)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(a)
	},
}

var appointmentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List appointments",
	Run: func(cmd *cobra.Command, args []string) {
		appmts, err := ListAppointments(appointment.ID, appointment.Status, db)
		if err != nil {
			log.Fatal(err)
		}

		for a := range appmts {
			fmt.Println(appmts[a])
		}
	},
}

var appointmentCompleteCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete an appointment and create a visit",
	Run: func(cmd *cobra.Command, args []string) {
		a, err := GetAppointment(appointment.ID, db)

		if !a.IsCreated() {
			log.Fatal(ErrAppointmentNotCreated)
		}

		v, err := CreateVisit(appointment.ID, visit.Notes, db)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(v)

		err = UpdateAppointmentStatus(appointment.ID, COMPLETED, db)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var appointmentCancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel an appointment",
	Run: func(cmd *cobra.Command, args []string) {
		a, err := GetAppointment(appointment.ID, db)

		if !a.IsCreated() {
			log.Fatal(ErrAppointmentNotCreated)
		}

		err = UpdateAppointmentStatus(appointment.ID, CANCELED, db)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var appointmentNoShowCmd = &cobra.Command{
	Use:   "noshow",
	Short: "Mark an appointment as 'noshow'",
	Run: func(cmd *cobra.Command, args []string) {
		a, err := GetAppointment(appointment.ID, db)

		if !a.IsCreated() {
			log.Fatal(ErrAppointmentNotCreated)
		}

		err = UpdateAppointmentStatus(appointment.ID, NO_SHOW, db)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	appointmentCmd.AddCommand(appointmentCreateCmd)
	appointmentCmd.AddCommand(appointmentListCmd)
	appointmentCmd.AddCommand(appointmentCancelCmd)
	appointmentCmd.AddCommand(appointmentCompleteCmd)
	appointmentCmd.AddCommand(appointmentNoShowCmd)

	//APPOINTMENTS
	appointmentCmd.PersistentFlags().Int64VarP(&appointment.PatientID, "patient-id", "p", 0, "id of the patient")

	// CREATE
	appointmentCreateCmd.Flags().Int64VarP(&appointment.PatientID, "patient-id", "p", 0, "id of the patient")
	appointmentCreateCmd.Flags().IntVarP(&appointment.DurationMinutes, "duration", "d", 30, "duration of the appointment (in minutes)")

	appointmentCreateCmd.MarkFlagRequired("patient-id")

	// COMPLETE
	appointmentCompleteCmd.Flags().StringVarP(&visit.Notes, "notes", "n", "", "notes of patient visit")
	appointmentCompleteCmd.Flags().Int64Var(&appointment.ID, "id", 0, "id of the patient")
	appointmentCompleteCmd.MarkFlagRequired("id")

	// CANCEL
	appointmentCancelCmd.PersistentFlags().Int64Var(&appointment.ID, "id", 0, "id of the patient")
	appointmentCancelCmd.MarkFlagRequired("id")

	// NOSHOW
	appointmentNoShowCmd.PersistentFlags().Int64Var(&appointment.ID, "id", 0, "id of the patient")
	appointmentNoShowCmd.MarkFlagRequired("id")
}
