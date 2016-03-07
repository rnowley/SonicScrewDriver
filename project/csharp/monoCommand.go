package csharp

import (
	"fmt"
	"strings"
)

// MonoCommand provides a representation of a call to the mono
// command.
type MonoCommand struct {
	CommandName    string
	ExecutableName string
	RunArguments   []string
}

// NewDefaultMonoCommand returns a MonoCommand with some default values set.
func NewDefaultMonoCommand() MonoCommand {
	var command MonoCommand
	command.CommandName = "mono"
	return command
}

// GetCommandName is a method which accesses the name of the command
// to be run.
func (command MonoCommand) GetCommandName() string {
	return command.CommandName
}

// GenerateArgumentList is a method which returns a slice of strings containing
// the arguments to use when running the mono command.
func (command MonoCommand) GenerateArgumentList() []string {
	argumentArray := make([]string, 0)

	argumentArray = append(argumentArray, command.ExecutableName)

	if len(command.RunArguments) != 0 {
		argumentArray = append(argumentArray, strings.Join(command.RunArguments, " "))
	}

	return argumentArray
}

func (command MonoCommand) String() string {
	arguments := strings.Join(command.GenerateArgumentList(), " ")
	return fmt.Sprintf("%s %s", command.GetCommandName(), arguments)
}
