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

var appointmentCmd = &cobra.Command{
	Use:   "appointment",
	Short: "Manage appointments",
}

var appointmentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an appointment",
	Run: func(cmd *cobra.Command, args []string) {
		if _, err := GetPatient(appointment.PatientID, db); err != nil {
			log.Fatal(err)
		}

		appointment.ScheduledFor = time.Now().UTC()

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
}

var appointmentCancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel an appointment",
}

var appointmentCompleteCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete an appointment and create a visit",
}

func init() {
	rootCmd.AddCommand(appointmentCmd)
	appointmentCmd.AddCommand(appointmentCreateCmd)
	appointmentCmd.AddCommand(appointmentListCmd)
	appointmentCmd.AddCommand(appointmentCancelCmd)
	appointmentCmd.AddCommand(appointmentCompleteCmd)

	appointmentCreateCmd.Flags().Int64VarP(&appointment.PatientID, "patient-id", "p", 0, "id of the patient")
	// appointmentCreateCmd.Flags().TimeVarP(&appointment.ScheduledFor, "schedule", "s", , "duration of the appointment (in minutes)")
	appointmentCreateCmd.Flags().IntVarP(&appointment.DurationMinutes, "duration", "d", 30, "duration of the appointment (in minutes)")
}
