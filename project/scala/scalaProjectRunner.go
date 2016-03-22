package scala

import (
	"bytes"
	"fmt"
	"os/exec"
)

type ScalaProjectRunner struct {
	command ScalaCommand
	project ScalaProject
}

func NewScalaProjectRunner(command ScalaCommand, project ScalaProject) ScalaProjectRunner {
	return ScalaProjectRunner{command, project}
}

func (runner ScalaProjectRunner) RunProject() error {
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
	// Attach buffer to command
	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdError

	// Execute command
	err := cmd.Run() // will wait for command to return
	printError(err)
	printOutput(cmdOutput.Bytes(), err != nil)
	printOutput(cmdError.Bytes(), err != nil)

	if err != nil {
		return err
	}

	return nil
}
