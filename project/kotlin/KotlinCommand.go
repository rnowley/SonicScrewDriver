package kotlin

import "strings"

// KotlinCommand contains all of the information required to
// call the Java command.
type KotlinCommand struct {
	ClassPath    []string
	CommandName  string
	JarFile      string
	MainClass    string
	RunArguments []string
}

// NewDefaultKotlinCommand returns a KotlinCommand with some default values
// set.
func NewDefaultKotlinCommand() KotlinCommand {
	var command KotlinCommand
	command.CommandName = "java"
	command.ClassPath = make([]string, 0, 10)
	command.JarFile = "./build/out.jar"

	return command
}

// GetCommandName is a method which accesses the name of the command to
// be run.
func (command KotlinCommand) GetCommandName() string {
	return command.CommandName
}

// GenerateArgumentList is a method that returns a slice of strings containing
// the arguments to use when running the Java command.
func (command KotlinCommand) GenerateArgumentList() []string {
	argumentArray := make([]string, 0)

	if len(command.ClassPath) != 0 {
		argumentArray = append(argumentArray, "-cp", strings.Join(command.ClassPath, ":"))
	}

	if command.JarFile != "" {
		argumentArray = append(argumentArray, "-jar", command.JarFile)
	} else {
		argumentArray = append(argumentArray, command.MainClass)
	}

	if len(command.RunArguments) != 0 {
		argumentArray = append(argumentArray, strings.Join(command.RunArguments, " "))
	}

	return argumentArray
}
