package cmd

import (
	"fmt"
	"os"
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
			os.Exit(1)
		}

		if err := appointment.IsValid(defaultWorkingHours); err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		overlap, err := CheckAppointmentOverlap(appointment.ScheduledFor, appointment.DurationMinutes, db)

		if overlap {
			appointments, err := ListAppointments(0, db)
			if err != nil {
				ui.Error(err)
				os.Exit(1)
			}

			suggestedScheduleFor, err := GetAvailableScheduleForAppointment(appointments, appointment.DurationMinutes, defaultWorkingHours, db)

			if err != nil {
				ui.Error(err)
				os.Exit(1)
			}

			if suggestedScheduleFor == "" {
				err = fmt.Errorf("the appointment for patient '%d' could not be scheduled: %v", appointment.PatientID, ErrAppointmentOverlap)
				ui.Error(err)
				os.Exit(1)
			}

			msg := fmt.Sprintf("Detected overlap with schedule time! There is an available schedule slot at: %s\nConfirm the change (y/N): ", suggestedScheduleFor)
			ui.Info(msg)

			var response string
			_, err = fmt.Scanln(&response)

			if err != nil || (response != "y" && response != "Y") {
				err = fmt.Errorf("Aborted")
				ui.Error(err)
				os.Exit(1)
			}

			appointment.ScheduledFor = suggestedScheduleFor
		}

		a, err := CreateAppointment(appointment, db)
		if err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		ui.Success("Appointment created")
		msg := fmt.Sprintf("Scheduled: %s", a.ScheduledFor)
		ui.Info(msg)
	},
}

var appointmentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List appointments",
	Run: func(cmd *cobra.Command, args []string) {
		appmts, err := ListAppointments(appointment.PatientID, db)
		if err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		if len(appmts) == 0 {
			if appointment.PatientID > 0 {
				ui.Info(fmt.Sprintf("no appointment found for patient %d", appointment.PatientID))
				os.Exit(0)
			}

			ui.Info("no appointment found")
		}
		var rows []ui.AppointmentRow

		for _, a := range appmts {
			var p Patient

			if a.PatientID != p.ID {
				p, err = GetPatient(a.PatientID, db)
				if err != nil {
					ui.Error(err)
					os.Exit(1)
				}
			}

			rows = append(rows, ui.AppointmentRow{
				ID:           a.ID,
				PatientName:  p.FullName(),
				ScheduledFor: a.ScheduledFor,
				Status:       string(a.Status),
			})
		}

		ui.RenderAppointments(rows)
	},
}

var appointmentUpdateCmd = &cobra.Command{
	Use:       "update",
	Short:     "Update an appointment status ['COMPLETE', 'CANCEL', 'NO_SHOW']",
	Args:      cobra.ExactArgs(1),
	ValidArgs: []string{"complete", "cancel", "noshow"},
}

var appointmentCompleteCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete an appointment and create a visit",
	Run: func(cmd *cobra.Command, args []string) {
		a, err := GetAppointment(appointment.ID, db)

		if !a.IsCreated() {
			ui.Error(ErrAppointmentNotCreated)
			os.Exit(1)
		}

		v, err := CreateVisit(appointment.ID, visit.Notes, db)
		if err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		err = UpdateAppointmentStatus(appointment.ID, COMPLETED, db)
		if err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		ui.Success("Appointment completed\n")
		ui.Info(fmt.Sprintf("Created visit with ID: %d", v.ID))
	},
}

var appointmentCancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel an appointment",
	Run: func(cmd *cobra.Command, args []string) {
		a, err := GetAppointment(appointment.ID, db)

		if !a.IsCreated() {
			ui.Error(ErrAppointmentNotCreated)
			os.Exit(1)
		}

		err = UpdateAppointmentStatus(appointment.ID, CANCELED, db)
		if err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		ui.Success("Appointment canceled")
	},
}

var appointmentNoShowCmd = &cobra.Command{
	Use:   "noshow",
	Short: "Mark an appointment as 'noshow'",
	Run: func(cmd *cobra.Command, args []string) {
		a, err := GetAppointment(appointment.ID, db)

		if !a.IsCreated() {
			ui.Error(ErrAppointmentNotCreated)
			os.Exit(1)
		}

		err = UpdateAppointmentStatus(appointment.ID, NO_SHOW, db)
		if err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		ui.Success("Appointment marked as noshow")
	},
}

func init() {
	appointmentCmd.AddCommand(appointmentCreateCmd)
	appointmentCmd.AddCommand(appointmentListCmd)
	appointmentCmd.AddCommand(appointmentUpdateCmd)
	appointmentUpdateCmd.AddCommand(appointmentCancelCmd, appointmentCompleteCmd, appointmentNoShowCmd)

	// CREATE
	appointmentCreateCmd.Flags().Int64VarP(&appointment.PatientID, "patient-id", "p", 0, "id of the patient")
	appointmentCreateCmd.Flags().IntVarP(&appointment.DurationMinutes, "duration", "d", 30, "duration of the appointment (in minutes)")
	appointmentCreateCmd.Flags().StringVarP(&appointment.ScheduledFor, "scheduled-for", "s", "", "The datetime <y-m-d hh:mm> (e.g.: '2026-04-16 10:30') the appointment will be scheduled")

	appointmentCreateCmd.MarkFlagRequired("patient-id")
	appointmentCreateCmd.MarkFlagRequired("scheduled-for")

	// UPDATE {"complete", "cancel", "noshow"}
	appointmentUpdateCmd.PersistentFlags().Int64Var(&appointment.ID, "id", 0, "id of the appointment")
	appointmentUpdateCmd.MarkPersistentFlagRequired("id")

	appointmentCompleteCmd.Flags().StringVarP(&visit.Notes, "notes", "n", "", "visit notes")
	appointmentCompleteCmd.MarkFlagRequired("notes")
}
