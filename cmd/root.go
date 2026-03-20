package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "smilelog",
	Short: "Smilelog is a CLI tool for dentists",
	Long: `Smilelog is a CLI tool designed to help dentists manage their work and day-to-day tasks.
			Documentation: ...
			Author: Bartolome Juan Des
			Repo: https://github.com/bartodes/smilelog`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("SmileLog CLI 🦷")
	},
}

func Execute() {
	rootCmd.Execute()
}
