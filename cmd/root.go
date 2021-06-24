package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "netcomp-cli",
		Short: "CLI tools for Netcomposer applications",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(encryptCmd, decryptCmd)
}
