package csharp

import (
	"bytes"
	"fmt"
	"github.com/rnowley/SonicScrewDriver/utilities"
	"os"
	"os/exec"
	"path/filepath"
)

// CSharpProjectBuilder represents a class for building a CSharp project.
type CSharpProjectBuilder struct {
	command CSharpCommand
	project CSharpProject
}

// New creates a new instance of a CSharpProjectBuilder.
func New(command CSharpCommand, project CSharpProject) CSharpProjectBuilder {
	return CSharpProjectBuilder{command, project}
}

// ExecutePreBuildTasks is used for executing any actions that need to be
// performed before building the project.
func (builder CSharpProjectBuilder) ExecutePreBuildTasks() error {
	err := builder.ensureDestinationDirectoryExists()
	return err
}

// BuildProject builds the Java project.
func (builder CSharpProjectBuilder) BuildProject() error {
	binary, lookErr := exec.LookPath(builder.command.GetCommandName())
	fmt.Println(binary)

	if lookErr != nil {
		return lookErr
	}

	args := builder.command.GenerateArgumentList()

	fmt.Println(args)
	cmd := exec.Command(binary, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	fmt.Printf("**%s\n", out.String())
	if err != nil {
		return err
	}

	return nil
}

// ExecutePostBuildTasks performs any tasks that need to be carried out after a
// successful build.
func (builder CSharpProjectBuilder) ExecutePostBuildTasks() error {
	fmt.Println("Post build tasks")
	builder.copyReferences()
	builder.copyResources()
	return nil
}

// copyReferences copies the required reference files to the build destination
// directory.
func (builder CSharpProjectBuilder) copyReferences() error {
	referenceCount := len(builder.project.References)
	fmt.Println(referenceCount)

	if referenceCount == 0 {
		return nil
	}

	destinationDirectory := builder.command.DestinationDirectory

	for i := 0; i < referenceCount; i++ {

		if builder.project.References[i].Path == "" {
			continue
		}

		path := builder.project.References[i].Path
		referenceName := builder.project.References[i].Name
		fileExtension := ".dll"

		utilities.CopyFile(fmt.Sprintf("%s%s%s", destinationDirectory, referenceName, fileExtension),
			fmt.Sprintf("%s%s%s", path, referenceName, fileExtension),
		)
	}
	return nil
}

// copyResources copies the required resources to the destination directory for the build.
func (builder CSharpProjectBuilder) copyResources() error {
	resourceCount := len(builder.project.Resources)

	if resourceCount == 0 {
		return nil
	}

	for i := 0; i < resourceCount; i++ {
		destinationDirectory := filepath.Dir(builder.project.Resources[i].Destination)

		utilities.EnsurePathExists(fmt.Sprintf("%s%s", builder.command.GetDestinationDirectory(), destinationDirectory))
		utilities.CopyFile(fmt.Sprintf("%s%s", builder.command.DestinationDirectory, builder.project.Resources[i].Destination), fmt.Sprintf("%s%s", builder.command.SourceDirectory, builder.project.Resources[i].Source))
	}

	return nil
}

// ensureDestinationDirectoryExists makes sure the the destination directory
// specified in the project already exists or if it doesn't then creates it.
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
