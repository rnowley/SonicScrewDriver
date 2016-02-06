package java

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type JavaProjectBuilder struct {
	command JavaCommand
	project JavaProject
}

func New(command JavaCommand, project JavaProject) JavaProjectBuilder {
	return JavaProjectBuilder{command, project}
}

func (builder JavaProjectBuilder) ExecutePreBuildTasks() error {
	err := builder.ensureDestinationDirectoryExists()
	return err
}

func (builder JavaProjectBuilder) BuildProject() error {
	binary, lookErr := exec.LookPath(builder.command.GetCommandName())

	if lookErr != nil {
		return lookErr
	}

	args := builder.command.GenerateArgumentList()
	fmt.Println(args)
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)

	if execErr != nil {
		return execErr
	}

	return nil
}

func (builder JavaProjectBuilder) ExecutePostBuildTasks() error {
	fmt.Println("Post build tasks (java)")
	return nil
}

func (builder JavaProjectBuilder) ensureDestinationDirectoryExists() error {
	destinationDirectory := builder.command.GetDestinationDirectory()
	_, err := os.Stat(destinationDirectory)

	if err != nil {
		err = os.MkdirAll(destinationDirectory, 0777)

		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("Created directory.")
		return nil
	}

	fmt.Println("File already exists, nothing to do.")
	return err
}
