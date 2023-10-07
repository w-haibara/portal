package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/w-haibara/portal/portal"
)

var rootCmd = &cobra.Command{
	Use:   "portal",
	Short: "",
	Long:  ``,
	RunE:  rootCmdRun,
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func rootCmdRun(cmd *cobra.Command, args []string) error {
	if err := portal.Run(); err != nil {
		return err
	}

	return nil
}
