package java

import (
	"fmt"
	"strings"
)

// JavacCommand provides a representation of a call to the Java
// compiler command.
type JavacCommand struct {
	CommandName          string
	SourceDirectory      string
	DestinationDirectory string
	SourceFiles          []string
	ClassPath            []string
	Deprecation          bool
	SourceVersion        string
	DebuggingInformation string
}

// NewDefaultJavacCommand returns a JavacCommand with some default values set.
func NewDefaultJavacCommand() JavacCommand {
	var command JavacCommand
	command.CommandName = "javac"
	command.SourceDirectory = "./src/"
	command.DestinationDirectory = "./build/"
	command.ClassPath = make([]string, 0, 10)
	return command
}

// GetCommandName is a method on a JavacCommand which accesses the name of the command
// to be run.
func (command JavacCommand) GetCommandName() string {
	return command.CommandName
}

// GetDestination is a method which returns the the destination directory where the
// compiler's output is going to be copied to.
func (command JavacCommand) GetDestinationDirectory() string {
	return command.DestinationDirectory
}

// GenerateArgumentList is a method which returns a slice of strings containing
// the arguments to use when running the java compiler command.
func (c JavacCommand) GenerateArgumentList() []string {
	argumentArray := make([]string, 0)
	argumentArray = append(argumentArray, "-d", c.DestinationDirectory)

	if c.DebuggingInformation != "" {
		argumentArray = append(argumentArray, c.DebuggingInformation)
	}

	if c.Deprecation {
		argumentArray = append(argumentArray, "-deprecation")
	}

	if len(c.SourceFiles) != 0 {
		argumentArray = append(argumentArray, c.SourceFiles...)
	}

	if len(c.ClassPath) != 0 {
		argumentArray = append(argumentArray, "-cp", strings.Join(c.ClassPath, ":"))
	}

	if c.SourceVersion != "" {
		argumentArray = append(argumentArray, "-source", c.SourceVersion)
	}

	return argumentArray
}

func (command JavacCommand) String() string {
	arguments := strings.Join(command.GenerateArgumentList(), " ")
	return fmt.Sprintf("%s %s", command.GetCommandName(), arguments)
}
