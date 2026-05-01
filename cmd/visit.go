package cmd

import (
	. "github.com/bartodes/smilelog/internals/models"
	. "github.com/bartodes/smilelog/internals/services"
	"github.com/bartodes/smilelog/internals/ui"
	"github.com/spf13/cobra"
)

var patientID int64

var visitCmd = &cobra.Command{
	Use:   "visit",
	Short: "see patient visits",
}

var listVisitCmd = &cobra.Command{
	Use:   "list",
	Short: "list patient visit",
	// Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := PatientExists(patientID, db); err != nil {
			ui.Error(err)
		}

		var visits []Visit
		visits, err := ListVisits(patientID, db)
		if err != nil {
			ui.Error(err)
		}

		if len(visits) == 0 {
			ui.Info("Patient has no visits")
		}

		var rows []ui.VisitRow

		for _, v := range visits {
			rows = append(rows, ui.VisitRow{
				ID:           v.ID,
				PatientName:  patient.Name,             //getPatient for name
				ScheduledFor: appointment.ScheduledFor, //getAppointment for ScheduledFor
				Notes:        v.Notes,
			})
		}

		ui.RenderVisits(rows)
	},
}

func init() {
	visitCmd.AddCommand(listVisitCmd)

	listVisitCmd.Flags().Int64VarP(&patientID, "patient-id", "p", 0, "the patient ID for whom you want to list the visits")
	listVisitCmd.MarkFlagRequired("patient-id")
}
