package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var patientCmd = &cobra.Command{
	Use:   "patient",
	Short: "Patient",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Executing patient...")
		return nil
	},
}

var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Patient",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Execution patient history...")
		return nil
	},
}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Patient",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Execution patient search...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(patientCmd)
	patientCmd.AddCommand(historyCmd)
	patientCmd.AddCommand(searchCmd)
}
