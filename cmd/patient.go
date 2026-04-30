package cmd

import (
	"fmt"
	"log"

	. "github.com/bartodes/smilelog/internals/models"
	. "github.com/bartodes/smilelog/internals/services"
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
			log.Fatal(err)
		}
		log.Println(p)
	},
}

var patientListCmd = &cobra.Command{
	Use:   "list",
	Short: "List patients",
	Run: func(cmd *cobra.Command, args []string) {
		patients, err := ListPatients(db)
		if err != nil {
			log.Fatal(err)
		}

		for _, patient := range patients {
			fmt.Println(patient)
		}
	},
}

var patientHistoryCmd = &cobra.Command{
	Use:   "history",
	Short: "Get a patient history summary",
	Run: func(cmd *cobra.Command, args []string) {
		if err := PatientExists(patient.ID, db); err != nil {
			log.Fatal(err)
		}

		appointments, err := ListAppointments(patient.ID, db)

		if err != nil {
			log.Fatal(err)
		}
		var as AppointmentStats

		for _, a := range appointments {
			as.Total++
			switch a.Status {
			case COMPLETED:
				as.Completed++
			case CANCELED:
				as.Canceled++
			case NO_SHOW:
				as.NoShow++
			}
		}

		visits, err := ListVisits(patient.ID, db)
		if err != nil {
			log.Fatal(err)
		}

		if len(appointments) > 0 {
			fmt.Printf("PATIENT HISTORY\nAPPOINTMENTS\n\tTotal: %d\n\tCompleted: %d\n\tCanceled: %d\n\tNoShow: %d\n",
				as.Total,
				as.Completed,
				as.Canceled,
				as.NoShow,
			)
			fmt.Printf("LAST APPOINTMENT: %+v\n", appointments[len(appointments)-1])
		}

		if len(visits) > 0 {
			fmt.Printf("LAST VISIT: %+v\n", visits[len(visits)-1])
		}

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

	patientHistoryCmd.Flags().Int64VarP(&patient.ID, "patient-id", "p", 0, "the id of the patient")
	patientHistoryCmd.MarkFlagRequired("patient-id")
}
