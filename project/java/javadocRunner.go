package java

import (
	"bytes"
	"fmt"
	"os/exec"
)

type JavadocBuilder struct {
	command JavadocCommand
	project JavaProject
}

func NewJavadocBuilder(command JavadocCommand, project JavaProject) JavadocBuilder {
	return JavadocBuilder{command, project}
}

func (builder JavadocBuilder) BuildDocumentation(verbose bool) error {
	binary, lookErr := exec.LookPath(builder.command.GetCommandName())

	if lookErr != nil {
		return lookErr
	}

	args := builder.command.GenerateArgumentList()

	if verbose {
		fmt.Println("Executing command:")
		fmt.Println(builder.command)
		fmt.Println("\nCompiler output:\n")
	}

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

	if verbose {
		fmt.Println("==========\n")
	}

	if err != nil {
		return err
	}

	return nil
}
