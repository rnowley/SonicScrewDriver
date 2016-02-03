package csharp

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
	argumentArray := make([]string, 0)
	argumentArray = append(argumentArray, c.CommandName)

	if c.DebugFlag != "" {
		argumentArray = append(argumentArray, c.DebugFlag)
	}

	if c.OutputFilename != "" {
		argumentArray = append(argumentArray, c.OutputFilename+GetFileSuffix(c.BuildTarget))
	}

	argumentArray = append(argumentArray, c.BuildTarget)

	if c.References != "" {
		argumentArray = append(argumentArray, c.References)
	}

	if len(c.SourceFiles) != 0 {
		argumentArray = append(argumentArray, c.SourceFiles...)
	}

	if len(c.SourceFiles) != 0 {
		argumentArray = append(argumentArray, c.LibraryPath)
	}

	if c.WarningLevel != "" {
		argumentArray = append(argumentArray, c.WarningLevel)
	}

	if c.WarningsAsErrors != "" {
		argumentArray = append(argumentArray, c.WarningsAsErrors)
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
