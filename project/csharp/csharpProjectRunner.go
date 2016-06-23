package csharp

import (
	//"bytes"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type CSharpProjectRunner struct {
	command MonoCommand
	project CSharpProject
}

func NewProjectRunner(command MonoCommand, project CSharpProject) CSharpProjectRunner {
	return CSharpProjectRunner{command, project}
}

func (runner CSharpProjectRunner) RunProject() error {
	binary, lookErr := exec.LookPath(runner.command.GetCommandName())

	if lookErr != nil {
		return lookErr
	}

	env := os.Environ()

	args := runner.command.GenerateArgumentList()
	fmt.Println(runner.command)

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
	panic(execErr)
	}

	return nil
}
