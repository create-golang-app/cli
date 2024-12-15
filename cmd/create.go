// Copyright 2024 Quy Duong and Create Go App Contributors. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.
package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolVarP(
		&useCustomTemplate,
		"template", "t", false,
		"enables to use custom backend and frontend templates",
	)
}

// createCmd represents the `create` command.
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new"},
	Short:   "Create a new project via interactive UI",
	Long:    "\nCreate a new project via interactive UI.",
	RunE:    runCreateCmd,
}

// runCreateCmd represents runner for the `create` command.
func runCreateCmd(cmd *cobra.Command, args []string) error {
	return nil
}
