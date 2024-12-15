package command

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"

	"github.com/create-golang-app/cli/pkg/utils/message"
)

// ExecCommand function to execute a given command.
func ExecCommand(command string, options []string, silentMode bool) error {
	if command == "" || options == nil {
		return fmt.Errorf("no command to execute")
	}

	// Create buffer for stderr.
	stderr := &bytes.Buffer{}

	// Collect command line.
	cmd := exec.Command(command, options...) // #nosec G204

	// Set buffer for stderr from cmd.
	cmd.Stderr = stderr

	// Create a new reader.
	cmdReader, errStdoutPipe := cmd.StdoutPipe()
	if errStdoutPipe != nil {
		return message.Error(errStdoutPipe.Error())
	}

	// Start executing command.
	if errStart := cmd.Start(); errStart != nil {
		return message.Error(errStart.Error())
	}

	// Create a new scanner and run goroutine func with output, if not in silent mode.
	if !silentMode {
		scanner := bufio.NewScanner(cmdReader)
		go func() {
			for scanner.Scan() {
				message.Info("", scanner.Text(), false, false)
			}
		}()
	}

	// Wait for executing command.
	if errWait := cmd.Wait(); errWait != nil {
		return message.Error(errWait.Error())
	}

	return nil
}
