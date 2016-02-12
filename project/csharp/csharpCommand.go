package csharp

import "strings"

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
}

// NewDefaultCommand returns a CSharpCommand with some default values set.
func NewDefaultCommand() CSharpCommand {
	var command CSharpCommand
	command.CommandName = "mcs"
	command.DebugFlag = "-debug"
	command.SourceDirectory = "./src/"
	command.DestinationDirectory = "./build"
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
	argumentArray := []string{
		c.CommandName,
		c.DebugFlag,
		c.OutputFilename + GetFileSuffix(c.BuildTarget),
		c.BuildTarget,
		c.LibraryPath,
		c.WarningLevel,
		c.WarningsAsErrors,
		c.References,
		strings.Join(c.SourceFiles, " "),
	}

	return argumentArray
}

// GetFileSuffix is a function for determining the file suffix of the build  artifact based in the provided build target.
func GetFileSuffix(buildTarget string) string {
	suffix := ".exe"

	switch buildTarget {
	case "target:exe":
		suffix = ".exe"
	case "target:library":
		suffix = ".dll"
	case "-target:module":
		suffix = ".netmodule"
	case "-target:winexe":
		suffix = ".exe"
	}

	return suffix
}
