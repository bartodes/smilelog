package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/bartodes/smilelog/internals/database"
	"github.com/spf13/cobra"
)

var db *sql.DB

var rootCmd = &cobra.Command{
	Use:   "smilelog",
	Short: "Smilelog is a CLI tool for dentists",
	Long: `Smilelog is a CLI tool designed to help dentists manage their work and day-to-day tasks.
			Documentation: ...
			Author: Bartolome Juan Des
			Repo: https://github.com/bartodes/smilelog`,
}

func init() {
	rootCmd.AddCommand(appointmentCmd)
	rootCmd.AddCommand(patientCmd)
	rootCmd.AddCommand(visitCmd)
}

func Execute() {
	db = database.InitDB()
	defer db.Close()

	// services.ListAppointmentsByScheduleRange()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
