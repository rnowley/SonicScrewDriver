package csharp

import (
	"fmt"
	"strings"
)
// CSharpCommand provides a representation of a call to the CSharp
// compiler command.
type CSharpCommand struct {
	CommandName          string
	DebugFlag            string
	OutputFilename       string
	SourceFiles          []string
	BuildTarget          string
	References           string
	SourceDirectory      string
	DestinationDirectory string
	LibraryPath          string
	PackageList          string
	WarningLevel         string
	WarningsAsErrors     string
	ReferencePaths       []Reference
}

// NewDefaultCommand returns a CSharpCommand with some default values set.
func NewDefaultCommand() CSharpCommand {
	var command CSharpCommand
	command.CommandName = "mcs"
	command.DebugFlag = "-debug"
	command.SourceDirectory = "./src/"
	command.DestinationDirectory = "./build/"
	return command
}

// GetCommandName is a method on a CSharpCommand which accesses the name of the command
// to be run.
func (command CSharpCommand) GetCommandName() string {
	return command.CommandName
}

// GetDestination is a method which returns the the destination directory where the
// compiler's output is going to be copied to.
func (command CSharpCommand) GetDestinationDirectory() string {
	return command.DestinationDirectory
}

// GenerateArgumentList is a method which returns a slice of strings containing
// the arguments to use when running the CSharp compiler command.
func (c CSharpCommand) GenerateArgumentList() []string {
	argumentArray := make([]string, 0)
	argumentArray = append(argumentArray, c.DebugFlag)
	argumentArray = append(argumentArray,
		c.OutputFilename+GetFileSuffix(c.BuildTarget))
	argumentArray = append(argumentArray, c.BuildTarget)

	if c.LibraryPath != "" {
		argumentArray = append(argumentArray, c.LibraryPath)
	}

	if c.WarningLevel != "" {
		argumentArray = append(argumentArray, c.WarningLevel)
	}

	if c.WarningsAsErrors != "" {
		argumentArray = append(argumentArray, c.WarningsAsErrors)
	}

	if c.References != "" {
		argumentArray = append(argumentArray, c.References)
	}

	if len(c.SourceFiles) != 0 {
		argumentArray = append(argumentArray, c.SourceFiles...)
	}

	return argumentArray
}

// GetFileSuffix is a function for determining the file suffix of the build  artifact based in the provided build target.
func GetFileSuffix(buildTarget string) string {
	var suffix string

	switch buildTarget {
	case "-target:exe":
		suffix = ".exe"
	case "-target:library":
		suffix = ".dll"
	case "-target:module":
		suffix = ".netmodule"
	case "-target:winexe":
		suffix = ".exe"
	default:
		suffix = ".exe"
	}

	return suffix
}

func (command CSharpCommand) String() string {
	arguments := strings.Join(command.GenerateArgumentList(), " ")
	return fmt.Sprintf("%s %s", command.GetCommandName(), arguments)
}
