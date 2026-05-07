package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/bartodes/smilelog/internals/database"
	"github.com/bartodes/smilelog/internals/ui"
	"github.com/spf13/cobra"
)

var db *sql.DB

var rootCmd = &cobra.Command{
	Use:   "smilelog",
	Short: "Smilelog is a CLI tool for dentists",
	Long:  "Smilelog is a CLI tool designed to help dentists manage their work and day-to-day tasks\n\nAuthor: Bartolome Juan Des\nDocumentation: https://github.com/bartodes/smilelog/blob/main/README.md\nRepo: https://github.com/bartodes/smilelog",
}

func init() {
	rootCmd.AddCommand(appointmentCmd)
	rootCmd.AddCommand(patientCmd)
	rootCmd.AddCommand(visitCmd)
}

func Execute() {
	var err error

	db, err = database.InitDB()
	if err != nil {
		ui.Error(err)
		os.Exit(1)
	}

	defer db.Close()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
