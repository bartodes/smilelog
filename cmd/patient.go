package cmd

import (
	"github.com/spf13/cobra"
)

var patientCmd = &cobra.Command{
	Use:   "patient",
	Short: "Manage patients",
}

var patientCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a atient",
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
	patientCmd.AddCommand(patientCmd)
	patientCmd.AddCommand(patientCreateCmd)
	patientCmd.AddCommand(patientListCmd)
	patientCmd.AddCommand(PatientHistoryCmd)
}
