package java

import (
	"strings"
)

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

func NewDefaultJavacCommand() JavacCommand {
	var command JavacCommand
	command.CommandName = "javac"
	command.SourceDirectory = "./src/"
	command.DestinationDirectory = "./build"
	command.ClassPath = make([]string, 0, 10)
	return command
}

func (command JavacCommand) GetCommandName() string {
	return command.CommandName
}

func (command JavacCommand) GetDestinationDirectory() string {
	return command.DestinationDirectory
}

func (c JavacCommand) GenerateArgumentList() []string {
	argumentArray := make([]string, 0)
	//argumentArray = append(argumentArray, c.CommandName)
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

