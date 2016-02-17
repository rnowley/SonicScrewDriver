package csharp

import (
	"bytes"
	"fmt"
	"os/exec"
)

type CSharpProjectRunner struct {
	command MonoCommand
	project CSharpProject
}

func NewProjectRunner(command MonoCommand, project CSharpProject) CSharpProjectRunner {
	return CSharpProjectRunner{command, project}
}

func (runner CSharpProjectRunner) RunProject() error {
	fmt.Println("Inside ProjectRunner.RunProject")
	binary, lookErr := exec.LookPath(runner.command.GetCommandName())
	fmt.Printf("Cmd to run: %s", binary)

	if lookErr != nil {
		return lookErr
	}

	args := runner.command.GenerateArgumentList()
	fmt.Println(args)

	// Create an *exec.Cmd
	cmd := exec.Command(binary, args...)

	// Stdout buffer
	cmdOutput := &bytes.Buffer{}
	cmdError := &bytes.Buffer{}
	// Attach buffer to command
	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdError

	// Execute command
	//printCommand(cmd)
	err := cmd.Run() // will wait for command to return
	printError(err)
	printOutput(cmdOutput.Bytes())
	printOutput(cmdError.Bytes())

	//cmd := exec.Command(binary, "")
	//cmd.Args = args
	/*var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr*/
	//err := cmd.Run()

	//f err != nil {
	//	return err
	//}

	/*fmt.Println(out.String())
	fmt.Println(stderr.String())*/

	return nil
}
