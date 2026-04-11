package cmd

import (
	"database/sql"
	"os"

	"github.com/bartodes/smilelog/internals/database"
	"github.com/spf13/cobra"
)

type App struct {
	DB *sql.DB
}

var rootCmd = &cobra.Command{
	Use:   "smilelog",
	Short: "Smilelog is a CLI tool for dentists",
	Long: `Smilelog is a CLI tool designed to help dentists manage their work and day-to-day tasks.
			Documentation: ...
			Author: Bartolome Juan Des
			Repo: https://github.com/bartodes/smilelog`,
}

func Execute() {
	db := database.InitDB()
	defer db.Close()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
