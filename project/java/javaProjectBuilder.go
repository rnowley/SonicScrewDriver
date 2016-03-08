package java

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
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
	err := builder.ensureDestinationDirectoryExists()
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
	fmt.Println(builder.command)

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

// ExecutePostBuildTasks performs any tasks that need to be carried out after a
// successful build.
func (builder JavaProjectBuilder) ExecutePostBuildTasks(verbose bool) error {
	return nil
}

// ensureDestinationDirectoryExists makes sure the the destination directory
// specified in the project already exists or if it doesn't then creates it.
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

	fmt.Println("Destination directory already exists, nothing to do.")
	return err
}
