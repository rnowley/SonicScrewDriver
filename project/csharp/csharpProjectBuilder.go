package csharp

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
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

// BuildProject builds the CSharp project.
func (builder CSharpProjectBuilder) BuildProject() error {
	binary, lookErr := exec.LookPath(builder.command.GetCommandName())

	if lookErr != nil {
		return lookErr
	}

	args := builder.command.GenerateArgumentList()
	fmt.Println(builder.command)

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

func printError(err error) {

	if err != nil {
		output := color.RedString("%s\n", err.Error())
		os.Stderr.WriteString(output)
	}

}

func printOutput(outs []byte, commandError bool) {

	if len(outs) > 0 {
		var output string

		if commandError {
			output = color.RedString("%s", string(outs))
		} else {
			output = color.YellowString("%s", string(outs))
		}

		fmt.Println(output)
	}

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

	fmt.Println("Destination directory already exists, nothing to do.")
	return err
}
