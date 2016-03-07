package kotlin

import (
	"bytes"
	"fmt"
	"os/exec"
)

type KotlinProjectRunner struct {
	command KotlinCommand
	project KotlinProject
}

// NewKotlinProjectRunner creates a new KotlinProjectRunner
func NewKotlinProjectRunner(command KotlinCommand, project KotlinProject) KotlinProjectRunner {
	return KotlinProjectRunner{command, project}
}

func (runner KotlinProjectRunner) RunProject() error {
	binary, lookErr := exec.LookPath(runner.command.GetCommandName())

	if lookErr != nil {
		return lookErr
	}

	args := runner.command.GenerateArgumentList()
	fmt.Println(runner.command)

	// Create an *exec.Cmd
	cmd := exec.Command(binary, args...)

	// Stdout buffer
	cmdOutput := &bytes.Buffer{}
	cmdError := &bytes.Buffer{}

	// Attach the buffers to the command.
	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdError

	// Execute the command.
	err := cmd.Run()
	printError(err)
	printOutput(cmdOutput.Bytes(), err != nil)
	printOutput(cmdError.Bytes(), err != nil)

	if err != nil {
		return err
	}

	return nil
}
