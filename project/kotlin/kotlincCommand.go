package kotlin

import (
	"strings"
)

// KotlincCommand provides a representation of a call to the Kotlin
// compiler command.
type KotlincCommand struct {
	BuildTarget          string
	ClassPath            []string
	CommandName          string
	DestinationDirectory string
	IncludeRuntime       bool
	SourceDirectory      string
	SourceFiles          []string
}

// NewDefaultKotlincCommand returns a KotlincCommand with some default values set.
func NewDefaultKotlincCommand() KotlincCommand {
	var command KotlincCommand
	command.ClassPath = make([]string, 0, 10)
	command.CommandName = "kotlinc"
	command.SourceDirectory = "./src/"
	command.SourceFiles = make([]string, 0, 10)

	return command
}

// GetCommandName is a method on a KotlincCommand which accesses the name of the command
// to be run.
func (command KotlincCommand) GetCommandName() string {
	return command.CommandName
}

// GenerateArgumentList is a method which returns a slice of strings containing
// the arguments to use when running the kotlinc compiler command.
func (command KotlincCommand) GenerateArgumentList() []string {
	argumentArray := make([]string, 0)
	argumentArray = append(argumentArray, "-d", command.DestinationDirectory)

	if len(command.ClassPath) != 0 {
		argumentArray = append(argumentArray, "-cp", strings.Join(command.ClassPath, ":"))
	}

	if command.BuildTarget == "executable" {
		argumentArray = append(argumentArray, "-include-runtime")
	}

	if len(command.SourceFiles) != 0 {
		argumentArray = append(argumentArray, command.SourceFiles...)
	}

	return argumentArray
}

// GetDestinationDirectory is a method which returns the the destination directory/jar where
// the compiler's output is going to be copied to
func (command KotlincCommand) GetDestinationDirectory() string {
	return command.DestinationDirectory
}
