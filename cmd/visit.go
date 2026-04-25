package cmd

import (
	"fmt"
	"log"

	. "github.com/bartodes/smilelog/internals/models"
	. "github.com/bartodes/smilelog/internals/services"
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
			log.Fatal(err)
		}

		var visits []Visit
		visits, err := ListVisits(patientID, db)
		if err != nil {
			log.Fatal(err)
		}

		if len(visits) == 0 {
			fmt.Println("patient has no visits")
		}

		for v := range visits {
			fmt.Println(visits[v])
		}
	},
}

func init() {
	visitCmd.AddCommand(listVisitCmd)

	listVisitCmd.Flags().Int64VarP(&patientID, "patient-id", "p", 0, "the patient ID for whom you want to list the visits")
	listVisitCmd.MarkFlagRequired("patient-id")
}
