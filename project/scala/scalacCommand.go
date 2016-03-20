package scala

import "strings"

// ScalacCommand provides all of the information to generate
// a call to the Scala compiler command.
type ScalacCommand struct {
	ClassPath            []string
	CommandName          string
	DebuggingInformation string
	Deprecation          bool
	DestinationDirectory string
	Encoding             string
	NoWarnings           bool
	Optimise             bool
	SourceDirectory      string
	SourceFiles          []string
	Target               string
	Verbose              bool
}

func NewDefaultScalacCommand() ScalacCommand {
	var command ScalacCommand
	command.CommandName = "scalac"
	command.SourceDirectory = "./src/"
	command.DestinationDirectory = "./build/"
	command.ClassPath = make([]string, 0, 10)
	return command
}

// GetCommandName is a method on a ScalacCommand which accesses the name of the command
// to be run.
func (command ScalacCommand) GetCommandName() string {
	return command.CommandName
}

// GetDestinationDirectory is a method on a ScalacCommand which accesses the Destination Directory command
// to be run.
func (command ScalacCommand) GetDestinationDirectory() string {
	return command.DestinationDirectory
}

// GenerateArgumentList is a method which returns a slice of strings containing
// the arguments to use when running the scalac compiler command.
func (command ScalacCommand) GenerateArgumentList() []string {
	argumentArray := make([]string, 0)
	argumentArray = append(argumentArray, "-d", command.DestinationDirectory)

	if command.Deprecation {
		argumentArray = append(argumentArray, "-deprecation")
	}

	if command.Verbose {
		argumentArray = append(argumentArray, "-verbose")
	}

	if command.Encoding != "" {
		argumentArray = append(argumentArray, "-encoding", command.Encoding)
	}

	if command.Target != "" {
		argumentArray = append(argumentArray, "-target", command.Target)
	}

	if command.Optimise {
		argumentArray = append(argumentArray, "-optimise")
	}

	if len(command.ClassPath) != 0 {
		argumentArray = append(argumentArray, "-classpath",
			strings.Join(command.ClassPath, ":"))
	}

	if command.DebuggingInformation != "" {
		argumentArray = append(argumentArray, command.DebuggingInformation)
	}

	if command.NoWarnings {
		argumentArray = append(argumentArray, "-nowarn")
	}

	if len(command.SourceFiles) != 0 {
		argumentArray = append(argumentArray, command.SourceFiles...)
	}

	return argumentArray
}
