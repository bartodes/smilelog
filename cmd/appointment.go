package cmd

import (
	"github.com/spf13/cobra"
)

var appointmentCmd = &cobra.Command{
	Use:   "appointment",
	Short: "Manage appointments",
}

var appointmentCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an appointment",
}

var appointmentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List appointments",
}

var appointmentCancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel an appointment",
}

var appointmentCompleteCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete an appointment and create a visit",
}

func init() {
	rootCmd.AddCommand(appointmentCmd)
	appointmentCmd.AddCommand(appointmentCreateCmd)
	appointmentCmd.AddCommand(appointmentListCmd)
	appointmentCmd.AddCommand(appointmentCancelCmd)
	appointmentCmd.AddCommand(appointmentCompleteCmd)
}
