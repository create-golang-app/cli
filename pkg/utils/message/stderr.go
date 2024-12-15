package message

import (
	"fmt"

	"github.com/create-golang-app/cli/pkg/utils"
)

func colorizeLevel(level string) string {
	// Define variables.
	var (
		red         = "\033[0;31m"
		green       = "\033[0;32m"
		yellow      = "\033[1;33m"
		noColor     = "\033[0m"
		color, icon string
	)

	// Switch color.
	switch level {
	case "success":
		color = green
		icon = "[OK]"
	case "error":
		color = red
		icon = "[ERROR]"
	case "info":
		color = yellow
		icon = "[INFO]"
	default:
		color = noColor
	}

	// Send common or colored caption.
	return fmt.Sprintf("%s%s%s", color, icon, noColor)
}

func Info(level, text string, startWithNewLine, endWithNewLine bool) {
	// Define variables.
	var startLine, endLine string

	if startWithNewLine {
		startLine = "\n" // set a new line
	}

	if endWithNewLine {
		endLine = "\n" // set a new line
	}

	// Formatting message.
	message := fmt.Sprintf("%s %s %s %s", startLine, colorizeLevel(level), text, endLine)

	// Return output.
	_, err := fmt.Fprintln(utils.Stdout, message)
	if err != nil {
		return
	}
}

func Error(e string) error {
	return fmt.Errorf("%s%s", colorizeLevel("error"), e)
}
