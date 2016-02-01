package project

import "strings"

func BuildCommand(configuration ProjectConfiguration, arguments Arguments) *Command {
	command := NewDefaultCommand()

	if configuration.DestinationDirectory != "" {
		command.DestinationDirectory = configuration.DestinationDirectory
	}

	if len(configuration.ClassPath) != 0 {
		command.ClassPath = configuration.ClassPath
	}

	command.Deprecation = arguments.Deprecation

	command.SourceFiles = ExtractSourceFileList(configuration, command.SourceDirectory)

	if configuration.SourceVersion != "" {
		command.SourceVersion = configuration.SourceVersion
	}

	command.DebuggingInformation = ExtractDebuggingInformation(configuration)

	return command
}

func ExtractSourceFileList(configuration ProjectConfiguration,
	sourceDirectory string) []string {
	fileCount := len(configuration.SourceFiles)
	fileList := make([]string, fileCount)

	for i := 0; i < fileCount; i++ {
		fileList[i] = sourceDirectory + configuration.SourceFiles[i]
	}

	return fileList
}

func ExtractDebuggingInformation(configuration ProjectConfiguration) string {

	if len(configuration.DebuggingInformation) == 0 {
		return ""
	}

	if len(configuration.DebuggingInformation) == 1 &&
		configuration.DebuggingInformation[0] == "all" {
		return "-g"
	}

	return "-g:" + strings.Join(configuration.DebuggingInformation, ",")
}
