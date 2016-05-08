package csharp

// CscProjectBuilder represents a class for building a .netcore CSharp project.

import (
	"bytes"
	"fmt"
	"github.com/rnowley/SonicScrewDriver/utilities"
	"os"
	"os/exec"
	"path/filepath"
)

// CscProjectBuilder represents a class for building a .netcore CSharp project.
type CscProjectBuilder struct {
	command CscCommand
	project CSharpProject
}

// New creates a new instance of a CscProjectBuilder.
func NewCscProjectBuilder(command CscCommand, project CSharpProject) CscProjectBuilder {
	return CscProjectBuilder{command, project}
}

// ExecutePreBuildTasks is used for executing any actions that need to be
// performed before building the project.
func (builder CscProjectBuilder) ExecutePreBuildTasks(verbose bool) error {

	if verbose {
		fmt.Println("==========")
		fmt.Println("Pre-build tasks")
		fmt.Println("==========\n")
	}

	err := builder.ensureDestinationDirectoryExists(verbose)

	if verbose {
		fmt.Println("\n==========")
	}

	return err
}

// BuildProject builds the .netcore CSharp project.
func (builder CscProjectBuilder) BuildProject(verbose bool) error {
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

// ExecutePostBuildTasks performs any tasks that need to be carried out after a
// successful build.
func (builder CscProjectBuilder) ExecutePostBuildTasks(verbose bool) error {

	if verbose {
		fmt.Println("==========")
		fmt.Println("Post-build tasks")
		fmt.Println("==========\n")
	}

	builder.copyReferences(verbose)
	builder.copyResources(verbose)

	if verbose {
		fmt.Println("\n==========")
	}

	return nil
}

// copyReferences copies the required reference files to the build destination
// directory.
func (builder CscProjectBuilder) copyReferences(verbose bool) error {

	if verbose {
		fmt.Println("----------")
		fmt.Println("Task: Copy references.")
		fmt.Println("----------\n")
	}

	referenceCount := len(builder.command.ReferencePaths)

	if referenceCount == 0 {
		return nil
	}

	destinationDirectory := builder.command.DestinationDirectory

	for i := 0; i < referenceCount; i++ {

		if builder.command.ReferencePaths[i].Path == "" {
			continue
		}

		path := builder.command.ReferencePaths[i].Path
		referenceName := builder.command.ReferencePaths[i].Name
		fileExtension := ".dll"

		utilities.CopyFile(fmt.Sprintf("%s%s%s", destinationDirectory, referenceName, fileExtension),
			fmt.Sprintf("%s%s%s", path, referenceName, fileExtension),
		)

		if verbose {
			fmt.Printf("Copying %s%s%s to %s%s%s\n", path, referenceName,
				fileExtension, destinationDirectory, referenceName,
				fileExtension)
		}
	}
	return nil
}

// copyResources copies the required resources to the destination directory for the build.
func (builder CscProjectBuilder) copyResources(verbose bool) error {

	if verbose {
		fmt.Println("----------")
		fmt.Println("Task: Copy resources.")
		fmt.Println("----------\n")
	}

	resourceCount := len(builder.project.Resources)

	if resourceCount == 0 {
		return nil
	}

	for i := 0; i < resourceCount; i++ {
		destinationDirectory := filepath.Dir(builder.project.Resources[i].Destination)

		utilities.EnsurePathExists(fmt.Sprintf("%s%s", builder.command.GetDestinationDirectory(), destinationDirectory))
		utilities.CopyFile(fmt.Sprintf("%s%s", builder.command.DestinationDirectory, builder.project.Resources[i].Destination), fmt.Sprintf("%s%s", builder.command.SourceDirectory, builder.project.Resources[i].Source))

		if verbose {
			fmt.Printf("Copying %s%s to %s%s", builder.command.SourceDirectory,
				builder.project.Resources[i].Source, builder.command.DestinationDirectory,
				builder.project.Resources[i].Destination)
		}
	}

	return nil
}

// ensureDestinationDirectoryExists makes sure the the destination directory
// specified in the project already exists or if it doesn't then creates it.
func (builder CscProjectBuilder) ensureDestinationDirectoryExists(verbose bool) error {
	destinationDirectory := builder.command.GetDestinationDirectory()
	_, err := os.Stat(destinationDirectory)

	if err != nil {
		err = os.MkdirAll(destinationDirectory, 0777)

		if err != nil {
			fmt.Println(err)
			return err
		}

		if verbose {
			fmt.Print("Created directory []%s.\n", destinationDirectory)
		}

		return nil
	}

	if verbose {
		fmt.Printf("Destination directory [%s] already exists, nothing to do.\n", destinationDirectory)
	}

	return err
}
