package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:     "ksp2-save-parser",
		Short:   "parser",
		Long:    ``,
		Version: "v0.1.0",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
