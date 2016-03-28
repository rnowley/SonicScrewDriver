package java

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/rnowley/SonicScrewDriver/utilities"
	"os"
	"os/exec"
	"path/filepath"
)

// JavaProjectBuilder represents a class for building a Java project.
type JavaProjectBuilder struct {
	command JavacCommand
	project JavaProject
}

// NewProjectBuilder creates a new instance of a JavaProjectBuilder.
func NewProjectBuilder(command JavacCommand, project JavaProject) JavaProjectBuilder {
	return JavaProjectBuilder{command, project}
}

// ExecutePreBuildTasks is used for executing any actions that need to be
// performed before building the project.
func (builder JavaProjectBuilder) ExecutePreBuildTasks(verbose bool) error {

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

// BuildProject builds the Java project.
func (builder JavaProjectBuilder) BuildProject(verbose bool) error {
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

// ExecutePostBuildTasks performs any tasks that need to be carried out after a
// successful build.
func (builder JavaProjectBuilder) ExecutePostBuildTasks(verbose bool) error {

	if verbose {
		fmt.Println("==========")
		fmt.Println("Post-build tasks")
		fmt.Println("==========\n")
	}

	builder.copyResources(verbose)

	if verbose {
		fmt.Println("\n==========")
	}

	return nil
}

// copyResources copies the required resources to the destination directory for the build.
func (builder JavaProjectBuilder) copyResources(verbose bool) error {

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
func (builder JavaProjectBuilder) ensureDestinationDirectoryExists(verbose bool) error {

	if verbose {
		fmt.Println("----------")
		fmt.Println("Task: Ensure destination directory exists.")
		fmt.Println("----------\n")
	}

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
