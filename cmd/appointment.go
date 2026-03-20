package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var apmntCmd = &cobra.Command{
	Use:     "appointment",
	Aliases: []string{"apmnt"},
	Short:   "Create an appointment",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Creating an appointment...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(apmntCmd)
}
