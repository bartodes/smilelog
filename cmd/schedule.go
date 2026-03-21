package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var scheduleCmd = &cobra.Command{
	Use:   "schedule",
	Short: "Schedule",
	// Args: ...,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Execution schedule...")
		return nil
	},
}

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Schedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Execution schedule today...")
		return nil
	},
}

var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "Schedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Execution schedule week...")
		return nil
	},
}

var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Schedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Execution schedule next...")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)
	scheduleCmd.AddCommand(todayCmd)
	scheduleCmd.AddCommand(weekCmd)
	scheduleCmd.AddCommand(nextCmd)
}
