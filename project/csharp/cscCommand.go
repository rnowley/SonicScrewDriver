package csharp

//import (
//	"fmt"
//	"strings"
//)

// CscCommand provides a representation of a call to the Microsoft .net core
// CSharp command.
type CscCommand struct {
	CommandName          string
	DebugFlag            string
	OutputFilename       string
	SourceFiles          []string
	BuildTarget          string
	References           string
	SourceDirectory      string
	DestinationDirectory string
	LibraryPath          string
	WarningLevel         string
	WarningsAsErrors     string
	ReferencePaths       []Reference
}

// NewDefaultCommand returns a CscCommand with some default values set.
func NewDefaultCscCommand() CscCommand {
	var command CscCommand
	command.CommandName = "csc"
	command.DebugFlag = "/debug"
	command.SourceDirectory = "./src/"
	command.DestinationDirectory = "./build/"
	return command
}

// GetCommandName is a method on the CscCommand which accesses the name
// of the command to be run.
func (command CscCommand) GetCommandName() string {
	return command.CommandName
}

// GetDestination is a method on a CscCommand which returns the
// destination directory where the compiler's output is going to
// be copied to.
func (command CscCommand) GetDestinationDirectory() string {
	return command.DestinationDirectory
}

func (command CscCommand) GenerateArgumentList() []string {
	argumentArray := make([]string, 0)
	argumentArray = append(argumentArray, command.DebugFlag)
	argumentArray = append(argumentArray,
		command.OutputFilename+GetFileSuffix(command.BuildTarget))
	argumentArray = append(argumentArray, command.BuildTarget)

	if command.LibraryPath != "" {
		argumentArray = append(argumentArray, command.LibraryPath)
	}

	if command.WarningLevel != "" {
		argumentArray = append(argumentArray, command.WarningLevel)
	}

	if command.WarningsAsErrors != "" {
		argumentArray = append(argumentArray, command.WarningsAsErrors)
	}

	if command.References != "" {
		argumentArray = append(argumentArray, command.References)
	}

	if len(command.SourceFiles) != 0 {
		argumentArray = append(argumentArray, command.SourceFiles...)
	}

	return argumentArray
}

// GetCscFileSuffix is a function for determining the file suffix of the build  artifact based in the provided build target.
func GetCscFileSuffix(buildTarget string) string {
	var suffix string

	switch buildTarget {
	case "/target:exe":
		suffix = ".exe"
	case "/target:library":
		suffix = ".dll"
	case "/target:module":
		suffix = ".netmodule"
	case "/target:winexe":
		suffix = ".exe"
	case "/target:winmdobj":
		suffix = ".winmdobj"
	default:
		suffix = ".exe"
	}

	return suffix
}
