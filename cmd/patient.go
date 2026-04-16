package cmd

import (
	"log"

	. "github.com/bartodes/smilelog/internals/models"
	"github.com/bartodes/smilelog/internals/services"
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
		p, err := services.CreatePatient(patient, db)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(p)
	},
}

var patientListCmd = &cobra.Command{
	Use:   "create",
	Short: "List patients",
}

var PatientHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Get a patient history summary",
}

func init() {
	rootCmd.AddCommand(patientCmd)
	patientCmd.AddCommand(patientCreateCmd)
	patientCmd.AddCommand(patientListCmd)
	patientCmd.AddCommand(PatientHistoryCmd)
	patientCreateCmd.Flags().StringVarP(&patient.Name, "name", "n", "", "name of the patient")
	patientCreateCmd.Flags().StringVarP(&patient.LastName, "lastname", "l", "", "last name of the patient")
	patientCreateCmd.Flags().StringVarP(&patient.Email, "email", "e", "", "email of the patient")
	patientCreateCmd.Flags().UintVarP(&patient.PhoneNumber, "phone", "p", 0, "phone number of the patitent")
}
