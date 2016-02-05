package csharp

import (
    "fmt"
    "os"
    "os/exec"
	"syscall"
)

type CSharpProjectBuilder struct {
    command CSharpCommand
    project CSharpProject
}

func New(command CSharpCommand, project CSharpProject) CSharpProjectBuilder {
    return CSharpProjectBuilder{command, project}
}

func (builder CSharpProjectBuilder) ExecutePreBuildTasks() error {
    err := builder.ensureDestinationDirectoryExists()
    return  err
}

func (builder CSharpProjectBuilder) BuildProject() error {
    return nil
}

func (builder CSharpProjectBuilder) ExecutePostBuildTasks() error {
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

func (builder CSharpProjectBuilder) copyReferences() error {
    referenceCount := len(builder.project.References)

    if referenceCount == 0 {
        return nil
    }

    //destinationDirectory := builder.command.DestinationDirectory

    for i := 0; i < referenceCount; i++ {

        if builder.project.References[i].Path == "" {
            continue
        }

        //path := builder.project.References[i].Path
        //referenceName := builder.project.References[i].Name
        //fileExtension := ".dll"
	}
    return nil
}

func (builder CSharpProjectBuilder) ensureDestinationDirectoryExists() error {
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