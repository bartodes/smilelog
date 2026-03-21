package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var appointmentCmd = &cobra.Command{
	Use:     "appointment",
	Aliases: []string{"apmnt"},
	Short:   "Create an appointment",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Creating an appointment...")
		return nil
	},
}

var bookCmd = &cobra.Command{
	Use:   "book",
	Short: "Appointment",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Execution appointment book...")
		return nil
	},
}

var reappointmentCmd = &cobra.Command{
	Use:   "reappointment",
	Short: "Appointment",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Execution appointment reappointment...")
		return nil
	},
}

var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Appointment",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Execution appointment cancel...")
		return nil
	},
}

var confirmCmd = &cobra.Command{
	Use:   "confirm",
	Short: "Appointment",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Execution appointment confirm...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(appointmentCmd)
	appointmentCmd.AddCommand(bookCmd)
	appointmentCmd.AddCommand(reappointmentCmd)
	appointmentCmd.AddCommand(cancelCmd)
	appointmentCmd.AddCommand(confirmCmd)
}
