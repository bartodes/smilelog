package cmd

import (
	"os"
	"strconv"

	. "github.com/bartodes/smilelog/internals/models"
	. "github.com/bartodes/smilelog/internals/services"
	"github.com/bartodes/smilelog/internals/ui"
	"github.com/spf13/cobra"
)

var patient Patient

var patientCmd = &cobra.Command{
	Use:   "patient",
	Short: "Manage patients",
}

var patientCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a patient",
	Run: func(cmd *cobra.Command, args []string) {
		p, err := CreatePatient(patient, db)
		if err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		ui.Success("Patient created")
		ui.Info("ID: " + strconv.Itoa(int(p.ID)))
		ui.Info("Name: " + p.Name)
		ui.Info("Last Name: " + p.LastName)
		ui.Info("Phone: " + strconv.Itoa(int(p.PhoneNumber)))
		ui.Info("Email: " + p.Email)
	},
}

var patientListCmd = &cobra.Command{
	Use:   "list",
	Short: "List patients",
	Run: func(cmd *cobra.Command, args []string) {
		patients, err := ListPatients(db)
		if err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		var rows []ui.PatientRow

		for _, p := range patients {
			rows = append(rows, ui.PatientRow{
				ID:          p.ID,
				Name:        p.FullName(),
				Email:       p.Email,
				PhoneNumber: p.PhoneNumber,
			})
		}

		ui.RenderPatients(rows)
	},
}

var patientHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Get a patient history summary",
	Run: func(cmd *cobra.Command, args []string) {
		if err := PatientExists(patient.ID, db); err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		p, err := GetPatient(patient.ID, db)
		if err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		appointments, err := ListAppointments(patient.ID, db)

		if err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		var appointmentRows []ui.AppointmentRow
		for _, a := range appointments {
			appointmentRows = append(appointmentRows, ui.AppointmentRow{
				ID:           a.ID,
				PatientName:  p.FullName(),
				ScheduledFor: a.ScheduledFor,
				Status:       string(a.Status),
			})
		}

		var stats AppointmentStats

		for _, a := range appointments {
			stats.Total++
			switch a.Status {
			case COMPLETED:
				stats.Completed++
			case CANCELED:
				stats.Canceled++
			case NO_SHOW:
				stats.NoShow++
			}
		}

		visits, err := ListVisits(patient.ID, db)
		if err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		var visitRows []ui.VisitRow
		for _, v := range visits {
			a, err := GetAppointment(v.AppointmentId, db)
			if err != nil {
				ui.Error(err)
				os.Exit(1)
			}

			visitRows = append(visitRows, ui.VisitRow{
				ID:           v.ID,
				PatientName:  p.FullName(),
				ScheduledFor: a.ScheduledFor,
				Notes:        v.Notes,
			})
		}

		historyView := ui.PatientHistoryView{
			PatientName: p.FullName(),

			TotalAppointments: stats.Total,
			Completed:         stats.Completed,
			Cancelled:         stats.Canceled,
			NoShow:            stats.NoShow,

			Appointments: appointmentRows,
			Visits:       visitRows,
		}

		ui.RenderPatientHistory(historyView)

	},
}

func init() {
	patientCmd.AddCommand(patientCreateCmd)
	patientCmd.AddCommand(patientListCmd)
	patientCmd.AddCommand(patientHistoryCmd)
	patientCreateCmd.Flags().StringVarP(&patient.Name, "name", "n", "", "name of the patient")
	patientCreateCmd.Flags().StringVarP(&patient.LastName, "lastname", "l", "", "last name of the patient")
	patientCreateCmd.Flags().StringVarP(&patient.Email, "email", "e", "", "email of the patient")
	patientCreateCmd.Flags().UintVarP(&patient.PhoneNumber, "phone", "p", 0, "phone number of the patitent")
	patientCreateCmd.MarkFlagRequired("name")
	patientCreateCmd.MarkFlagsOneRequired("email", "phone")

	patientHistoryCmd.Flags().Int64VarP(&patient.ID, "patient-id", "p", 0, "the id of the patient")
	patientHistoryCmd.MarkFlagRequired("patient-id")
}
