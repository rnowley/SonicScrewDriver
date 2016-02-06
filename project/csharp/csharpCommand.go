package csharp

import "strings"

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

func NewDefaultCommand() CSharpCommand {
	var command CSharpCommand
	command.CommandName = "mcs"
	command.DebugFlag = "-debug"
	command.SourceDirectory = "./src/"
	command.DestinationDirectory = "./build"
	return command
}

func (command CSharpCommand) GetCommandName() string {
	return command.CommandName
}

func (command CSharpCommand) GetDestinationDirectory() string {
	return command.DestinationDirectory
}

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
