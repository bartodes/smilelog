package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var visitCmd = &cobra.Command{
	Use:   "visit",
	Short: "Visit",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Execution visit...")
		return nil
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Visit",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Execution visit add...")
		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Visit",
	RunE: func(cmd *cobra.Command, args []string) error {
		patient, err := cmd.Flags().GetInt("patient")
		if err != nil {
			return err
		}
		fmt.Printf("Execution visit list for %d\n", patient)
		return nil
	},
}

var notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "Visit",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Execution visit notes...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(visitCmd)
	visitCmd.AddCommand(addCmd)
	listCmd.Flags().IntP("patient", "p", 12345678, "list visits by patient id")
	visitCmd.AddCommand(listCmd)
	visitCmd.AddCommand(notesCmd)
}
