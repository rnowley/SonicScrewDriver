package java

import (
	"strings"
)

// GetJavaBuildCommand is a function for building up a javac command that can be used for building
// a java project. This command is built up using the project configuration and the command line
// arguments passed in.
func GetJavaBuildCommand(configuration JavaProject, deprecation bool, verbose bool) JavacCommand {
	command := NewDefaultJavacCommand()

	if configuration.DestinationDirectory != "" {
		command.DestinationDirectory = configuration.DestinationDirectory
	}

	if configuration.SourceDirectory != "" {
		command.SourceDirectory = configuration.SourceDirectory
	}

	if len(configuration.ClassPath) != 0 {
		command.ClassPath = configuration.ClassPath
	}

	command.Deprecation = deprecation

	command.SourceFiles = ExtractSourceFileList(configuration, command.SourceDirectory)

	if configuration.SourceVersion != "" {
		command.SourceVersion = configuration.SourceVersion
	}

	command.Encoding = configuration.Encoding

	command.DebuggingInformation = ExtractDebuggingInformation(configuration)

	command.LintWarnings = ExtractLintWarnings(configuration)

	command.Verbose = verbose

	command.Target = configuration.Target

	return command
}

// ExtractSourceFileList is a function that reads all of the source files to be
// compiled from the configuration file and returns a slice of source files to be
// compiled using the javac command. Each source file has had the base path appended
// to it when returned from the function.
func ExtractSourceFileList(configuration JavaProject,
	sourceDirectory string) []string {
	fileCount := len(configuration.SourceFiles)
	fileList := make([]string, fileCount)

	for i := 0; i < fileCount; i++ {
		fileList[i] = sourceDirectory + configuration.SourceFiles[i]
	}

	return fileList
}

// ExtractDebuggingInformation is a function that builds up the debugging information
// flag for the compiler to determine what debugging information needs to generated
// with the compiled classes.
func ExtractDebuggingInformation(configuration JavaProject) string {

	if len(configuration.DebuggingInformation) == 0 {
		return ""
	}

	if len(configuration.DebuggingInformation) == 1 &&
		configuration.DebuggingInformation[0] == "all" {
		return "-g"
	}

	return "-g:" + strings.Join(configuration.DebuggingInformation, ",")
}

func ExtractLintWarnings(configuration JavaProject) string {

	if len(configuration.LintWarnings) == 0 {
		return ""
	}

	if len(configuration.LintWarnings) == 1 &&
		configuration.LintWarnings[0] == "all" {
		return "-Xlint"
	}

	return "-Xlint:" + strings.Join(configuration.LintWarnings, ",")
}
