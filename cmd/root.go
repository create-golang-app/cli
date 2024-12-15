// Copyright 2024 Quy Duong and Create Go App Contributors. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.
package cmd

import (
	"github.com/create-golang-app/cli/pkg/registry"
	"github.com/create-golang-app/cli/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	useCustomTemplate bool // define custom templates
)

var rootCmd = &cobra.Command{
	Use:     "cgapp",
	Version: registry.CLIVersion,
	Short:   "A powerful CLI for the Create Go App project",
	Long: `
A powerful CLI for the Create Go App project.

Create a new production-ready project with backend (Golang),
frontend (JavaScript, TypeScript) and deploy automation
(Ansible, Docker) by running one CLI command.

-> Focus on writing code and thinking of business logic!
<- The Create Go App CLI will take care of the rest.

A helpful documentation and next steps -> https://github.com/create-golang-app/cli/wiki`,
}

func init() {
	rootCmd.SetErr(utils.Stderr)
	rootCmd.SetOut(utils.Stdout)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	_ = rootCmd.Execute()
}
