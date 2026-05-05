package cmd

import (
	"fmt"
	"os"

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
			os.Exit(1)
		}

		p, err := GetPatient(patientID, db)
		if err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		visits, err := ListVisits(patientID, db)
		if err != nil {
			ui.Error(err)
			os.Exit(1)
		}

		if len(visits) == 0 {
			ui.Info(fmt.Sprintf("no visit found for patient %d", patientID))
			os.Exit(0)
		}

		var rows []ui.VisitRow

		for _, v := range visits {
			a, err := GetAppointment(v.AppointmentId, db)
			if err != nil {
				ui.Error(err)
				os.Exit(1)
			}

			rows = append(rows, ui.VisitRow{
				ID:           v.ID,
				PatientName:  p.FullName(),
				ScheduledFor: a.ScheduledFor,
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
