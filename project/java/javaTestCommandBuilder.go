package java

import (
	"github.com/rnowley/SonicScrewDriver/project"
	"strings"
)

func BuildTestCommand(configuration JavaProject, arguments project.Arguments) JavacCommand {
	command := NewDefaultJavacCommand()

	command.DestinationDirectory = configuration.TestProject.DestinationDirectory

	classPathCount := len(configuration.ClassPath)

	// Add the current directory to the command's class path.
	command.ClassPath = append(command.ClassPath, ".")

	// Add the project's class path to the command's class path.
	for i := 0; i < classPathCount; i++ {
		command.ClassPath = append(command.ClassPath, configuration.ClassPath[i])
	}

	classPathCount = len(configuration.TestProject.ClassPath)

	// Add the test project's class path to the command's class path
	for i := 0; i < classPathCount; i++ {
		command.ClassPath = append(command.ClassPath, configuration.TestProject.ClassPath[i])
	}

	command.Deprecation = arguments.Deprecation

	command.SourceFiles = ExtractTestSourceFileList(configuration, command.SourceDirectory)

	if configuration.SourceVersion != "" {
		command.SourceVersion = configuration.SourceVersion
	}

	command.DebuggingInformation = ExtractDebuggingInformation(configuration)

	return command
}

func ExtractTestSourceFileList(configuration JavaProject,
	sourceDirectory string) []string {
	fileCount := len(configuration.TestProject.SourceFiles)
	fileList := make([]string, fileCount)

	for i := 0; i < fileCount; i++ {
		fileList[i] = sourceDirectory + configuration.TestProject.SourceFiles[i]
	}

	return fileList
}

func ExtractTestDebuggingInformation(configuration JavaProject) string {

	if len(configuration.DebuggingInformation) == 0 {
		return ""
	}

	if len(configuration.DebuggingInformation) == 1 &&
		configuration.DebuggingInformation[0] == "all" {
		return "-g"
	}

	return "-g:" + strings.Join(configuration.DebuggingInformation, ",")
}
