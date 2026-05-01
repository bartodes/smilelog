package cmd

import (
	"fmt"
	"log"
	"time"

	. "github.com/bartodes/smilelog/internals/models"
	. "github.com/bartodes/smilelog/internals/services"
	"github.com/bartodes/smilelog/internals/ui"
	"github.com/spf13/cobra"
)

var appointment Appointment
var visit Visit
var defaultWorkingHours = WorkingSchedule{
	Days: map[time.Weekday]bool{
		time.Monday:    true,
		time.Tuesday:   true,
		time.Wednesday: true,
		time.Thursday:  true,
		time.Friday:    true,
	},
	Start: 8,
	End:   18,
}

var appointmentCmd = &cobra.Command{
	Use:   "appointment",
	Short: "Manage appointments",
}

var appointmentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an appointment",
	Run: func(cmd *cobra.Command, args []string) {
		if err := PatientExists(appointment.PatientID, db); err != nil {
			ui.Error(err)
		}

		if ok, err := appointment.IsValid(defaultWorkingHours); err != nil {
			ui.Error(err)
		} else if !ok {
			log.Fatal(ErrInvalidAppointment)
		}

		overlap, err := CheckAppointmentOverlap(appointment.ScheduledFor, appointment.DurationMinutes, db)

		if overlap {
			appointments, err := ListAppointments(0, db)
			if err != nil {
				ui.Error(err)
			}

			suggestedScheduleFor, err := GetAvailableScheduleForAppointment(appointments, appointment.DurationMinutes, defaultWorkingHours, db)

			if err != nil {
				ui.Error(err)
			}

			if suggestedScheduleFor == "" {
				log.Fatalf("the appointment for patient '%d' could not be scheduled: %v", appointment.PatientID, ErrAppointmentOverlap)
			}

			fmt.Printf("Detected overlap with schedule time! There is an available schedule slot at: %s\nConfirm the change (y/N): ", suggestedScheduleFor)

			var response string
			_, err = fmt.Scanln(&response)

			if err != nil || (response != "y" && response != "Y") {
				log.Fatal("Aborted")
				return
			}

			appointment.ScheduledFor = suggestedScheduleFor
		}

		a, err := CreateAppointment(appointment, db)
		if err != nil {
			ui.Error(err)
		}

		ui.Success("Appointment created")
		ui.Info(fmt.Sprintf("Scheduled: %s", a.ScheduledFor))
	},
}

var appointmentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List appointments",
	Run: func(cmd *cobra.Command, args []string) {
		appmts, err := ListAppointments(appointment.PatientID, db)
		if err != nil {
			ui.Error(err)
		}

		if appointment.PatientID > 0 {
			fmt.Printf("Listing appointments of patient: %d\n", appointment.PatientID)
		}
		var rows []ui.AppointmentRow

		for _, a := range appmts {
			rows = append(rows, ui.AppointmentRow{
				ID:           a.ID,
				ScheduledFor: a.ScheduledFor,
				Status:       string(a.Status),
			})
		}

		ui.RenderAppointments(rows)
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
			ui.Error(err)
		}

		fmt.Println(v)

		err = UpdateAppointmentStatus(appointment.ID, COMPLETED, db)
		if err != nil {
			ui.Error(err)
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
			ui.Error(err)
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
			ui.Error(err)
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
	appointmentCreateCmd.Flags().StringVarP(&appointment.ScheduledFor, "scheduled-for", "s", "", "The datetime <y-m-d hh:mm> (e.g.: '2026-04-16 10:30') the appointment will be scheduled")

	appointmentCreateCmd.MarkFlagRequired("patient-id")
	appointmentCreateCmd.MarkFlagRequired("scheduled-for")

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
