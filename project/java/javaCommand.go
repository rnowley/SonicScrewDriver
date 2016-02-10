package java

import (
	"strings"
)

type JavaCommand struct {
	CommandName  string
	ClassPath    []string
	JarFile      string
	MainClass    string
	RunArguments []string
}

func NewDefaultJavaCommand() JavaCommand {
	var command JavaCommand
	command.CommandName = "java"
	command.ClassPath = make([]string, 0, 10)
	return command
}

func (command JavaCommand) GetCommandName() string {
	return command.CommandName
}

func (c JavaCommand) GenerateArgumentList() []string {
	argumentArray := make([]string, 0)
	//argumentArray = append(argumentArray, c.CommandName)

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

