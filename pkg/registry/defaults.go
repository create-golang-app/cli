package registry

import (
	"github.com/AlecAivazis/survey/v2"
)

// CLIVersion version of Create Go App CLI.
const CLIVersion string = "1.0.0"

// CreateAnswers struct for a survey's answers for `create` command.
type CreateAnswers struct {
	Backend       string
	Frontend      string
	Proxy         string
	AgreeCreation bool `survey:"agree"`
}

var (
	// CreateQuestions survey's questions for `create` command.
	CreateQuestions = []*survey.Question{
		{
			Name: "backend",
			Prompt: &survey.Select{
				Message: "Choose a backend framework:",
				Options: []string{
					"none",
					"nodejs-typescript-express",
					"nodejs-typescript-fastify",
					"nodejs-typescript-fiber",
					"go-gin",
					"go-fiber",
					"go-echo",
					"go-zero",
					"go-kratos",
					"python-fastapi",
					"python-django",
					"python-flask",
					"java-springboot",
					"java-spring",
					"aspnetcore",
				},
				Default:  "go-echo",
				PageSize: 3,
			},
			Validate: survey.Required,
		},
		{
			Name: "frontend",
			Prompt: &survey.Select{
				Message: "Choose a frontend framework/library:",
				Help:    "Option with a `*-ts` tail will create a TypeScript template.",
				Options: []string{
					"none",
					"vanilla",
					"vanilla-ts",
					"react",
					"react-ts",
					"react-swc",
					"react-swc-ts",
					"preact",
					"preact-ts",
					"next",
					"next-ts",
					"nuxt",
					"vue",
					"vue-ts",
					"sveltekit",
					"svelte",
					"svelte-ts",
					"solid",
					"solid-ts",
					"lit",
					"lit-ts",
					"qwik",
					"qwik-ts",
				},
				Default:  "none",
				PageSize: 21,
			},
		},
		{
			Name: "proxy",
			Prompt: &survey.Select{
				Message: "Choose a web/proxy server:",
				Options: []string{
					"none",
					"traefik",
					"traefik-acme-dns",
					"nginx",
				},
				Default:  "none",
				PageSize: 4,
			},
		},
		{
			Name: "agree",
			Prompt: &survey.Confirm{
				Message: "If everything is okay, can I create this project for you? ;)",
				Default: true,
			},
		},
	}

	// CustomCreateQuestions survey's questions for `create -c` command.
	CustomCreateQuestions = []*survey.Question{
		{
			Name: "backend",
			Prompt: &survey.Input{
				Message: "Enter URL to the custom backend repository:",
			},
			Validate: survey.Required,
		},
		{
			Name: "frontend",
			Prompt: &survey.Input{
				Message: "Enter URL to the custom frontend repository:",
				Default: "none",
			},
		},
		{
			Name: "proxy",
			Prompt: &survey.Select{
				Message: "Choose a web/proxy server:",
				Options: []string{
					"none",
					"traefik",
					"traefik-acme-dns",
					"nginx",
				},
				Default:  "none",
				PageSize: 4,
			},
		},
		{
			Name: "agree",
			Prompt: &survey.Confirm{
				Message: "If everything is okay, can I create this project for you? ;)",
				Default: true,
			},
		},
	}
)
