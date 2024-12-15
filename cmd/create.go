// Copyright 2024 Quy Duong and Create Go App Contributors. All rights reserved.
// Use of this source code is governed by Apache 2.0 license
// that can be found in the LICENSE file.
package cmd

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/create-golang-app/cli/pkg/command"
	"github.com/create-golang-app/cli/pkg/git"
	"github.com/create-golang-app/cli/pkg/registry"
	"github.com/create-golang-app/cli/pkg/utils/message"
	"github.com/spf13/cobra"
)

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
	if useCustomTemplate {
		// Ask with survey
		if err := survey.Ask(
			registry.CustomCreateQuestions,
			&createAnswers,
			survey.WithIcons(surveyIconsConfig),
		); err != nil {
			return message.Error(err.Error())
		}

		backend = createAnswers.Backend
	} else {
		// Ask with survey
		if err := survey.Ask(
			registry.CreateQuestions,
			&createAnswers,
			survey.WithIcons(surveyIconsConfig),
		); err != nil {
			return message.Error(err.Error())
		}

		backend = fmt.Sprintf(
			"github.com/create-golang-app/%v-template",
			strings.ReplaceAll(createAnswers.Backend, "/", "_"),
		)
	}

	if !createAnswers.AgreeCreation {
		message.Info(
			"info",
			"Oh no! You said \"no\", so I won't create anything. Hope to see you soon!",
			true,
			true,
		)

		return nil
	}

	_ = CloneBackend(backend)

	// Proccess bar for create project.
	if createAnswers.Frontend != "none" {
		_ = CloneFrontend()
	}

	message.Info(
		"",
		"Have a happy new project! :)",
		false, true,
	)
	return nil
}

func CloneBackend(backendUrl string) error {
	if err := git.GitClone("backend", backendUrl); err != nil {
		return message.Error(err.Error())
	}

	message.Info(
		"success",
		fmt.Sprintf("Backend was created with template `%v`!", backendUrl),
		true,
		false,
	)

	return nil
}

func CloneFrontend() error {
	if useCustomTemplate {
		if err := git.GitClone("frontend", createAnswers.Frontend); err != nil {
			return message.Error(err.Error())
		}

		message.Info(
			"success",
			fmt.Sprintf("Frontend was created with template `%v`!", createAnswers.Frontend),
			false,
			false,
		)
		return nil
	}

	switch createAnswers.Frontend {
	case "next":
		// Create a default frontend template with Next.js (React).
		if err := command.ExecCommand(
			"npx",
			[]string{
				"create-next-app@latest", "frontend",
				"--javascript",
				"--eslint",
				"--app",
				"--tailwind", "false",
				"--src-dir", "false",
				"--import-alias", "false",
			}, true,
		); err != nil {
			return err
		}
	case "next-ts":
		// Create a default frontend template with Next.js (React, Typescript).
		if err := command.ExecCommand(
			"npx",
			[]string{
				"create-next-app@latest", "frontend",
				"--typescript",
				"--eslint",
				"--app",
				"--tailwind", "false",
				"--src-dir", "false",
				"--import-alias", "false",
			}, true,
		); err != nil {
			return err
		}
	default:
		// Create a default frontend template from Vite (Pure JS/TS, React, Preact, Vue, Svelte, Lit).
		if err := command.ExecCommand(
			"npm",
			[]string{
				"create", "vite@latest", "frontend",
				"--",
				"--template",
				createAnswers.Frontend,
			}, true,
		); err != nil {
			return err
		}
	}

	message.Info(
		"success",
		fmt.Sprintf("Frontend was created with template `%v`!", createAnswers.Frontend),
		false,
		false,
	)
	return nil
}
