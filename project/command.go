package project

import "strings"

type Command struct {
	CommandName          string
	SourceDirectory      string
	DestinationDirectory string
	SourceFiles          []string
	ClassPath            []string
	Deprecation          bool
	SourceVersion        string
	DebuggingInformation string
}

func NewDefaultCommand() *Command {
	command := new(Command)
	command.CommandName = "javac"
	command.SourceDirectory = "./src/"
	command.DestinationDirectory = "./build"
	return command
}

func (c *Command) GenerateArgumentList() []string {
	argumentArray := make([]string, 0)
	argumentArray = append(argumentArray, c.CommandName)
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
