package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var patientCmd = &cobra.Command{
	Use:   "patient",
	Short: "Find or create a patient",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Finding or creating patient...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(patientCmd)
}
