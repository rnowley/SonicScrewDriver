package scala

import (
	"fmt"
	"strings"
)

// ScalaCommand provides a representation of a call to the Scala
// command.
type ScalaCommand struct {
	CommandName  string
	ClassPath    []string
	JarFile      string
	MainClass    string
	RunArguments []string
}

// NewDefaultScalaCommand returns a ScalaCommand with some default values set.
func NewDefaultScalaCommand() ScalaCommand {
	var command ScalaCommand
	command.CommandName = "scala"
	command.ClassPath = make([]string, 0, 10)
	return command
}

// GetCommandName is a method which accesses the name of the command
// to be run.
func (command ScalaCommand) GetCommandName() string {
	return command.CommandName
}

// GenerateArgumentList is a method which returns a slice of strings containing
// the arguments to use when running the scala command.
func (command ScalaCommand) GenerateArgumentList() []string {
	argumentArray := make([]string, 0)

	if len(command.ClassPath) != 0 {
		argumentArray = append(argumentArray, "-classpath", strings.Join(command.ClassPath, ":"))
	}

	if command.JarFile != "" {
		argumentArray = append(argumentArray, "-jar", command.JarFile)
	}

	if command.JarFile == "" && command.MainClass != "" {
		argumentArray = append(argumentArray, command.MainClass)
	}

	if len(command.RunArguments) != 0 {
		argumentArray = append(argumentArray, strings.Join(command.RunArguments, " "))
	}

	return argumentArray
}

func (command ScalaCommand) String() string {
	arguments := strings.Join(command.GenerateArgumentList(), " ")
	return fmt.Sprintf("%s %s", command.GetCommandName(), arguments)
}
