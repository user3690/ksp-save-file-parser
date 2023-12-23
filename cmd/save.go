package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/user3690/ksp-save-file-parser/src/saveparser"
)

var (
	saveCmd = &cobra.Command{
		Use:   "save",
		Short: "Parses the ksp save file",
		Long:  ``,
		RunE:  parse,
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.ExactArgs(1)(cmd, args); err != nil {
				return fmt.Errorf("missing path to save file: %w", err)
			}

			return nil
		},
	}
)

func init() {
	saveCmd.Flags().BoolP("saveToFile", "s", false, "save output to reports.txt")
}

func parse(cmd *cobra.Command, args []string) error {
	var (
		saveToFile bool
		err        error
	)

	saveToFile, err = cmd.Flags().GetBool("saveToFile")
	if err != nil {
		return err
	}

	err = saveparser.Parse(args[0], saveToFile)
	if err != nil {
		return err
	}

	return nil
}
