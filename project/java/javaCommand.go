package java

import (
	"strings"
)

// JavaCommand provides a representation of a call to the Java
// command.
type JavaCommand struct {
	CommandName  string
	ClassPath    []string
	JarFile      string
	MainClass    string
	RunArguments []string
}

// NewDefaultJavaCommand returns a JavaCommand with some default values set.
func NewDefaultJavaCommand() JavaCommand {
	var command JavaCommand
	command.CommandName = "java"
	command.ClassPath = make([]string, 0, 10)
	return command
}

// GetCommandName is a method which accesses the name of the command
// to be run.
func (command JavaCommand) GetCommandName() string {
	return command.CommandName
}

// GenerateArgumentList is a method which returns a slice of strings containing
// the arguments to use when running the java command.
func (c JavaCommand) GenerateArgumentList() []string {
	argumentArray := make([]string, 0)

	if len(c.ClassPath) != 0 {
		argumentArray = append(argumentArray, "-cp", strings.Join(c.ClassPath, ":"))
	}

	if c.JarFile != "" {
		argumentArray = append(argumentArray, "-jar", c.JarFile)
	}

	if c.JarFile == "" && c.MainClass != "" {
		argumentArray = append(argumentArray, c.MainClass)
	}

	if len(c.RunArguments) != 0 {
		argumentArray = append(argumentArray, strings.Join(c.RunArguments, " "))
	}

	return argumentArray
}
