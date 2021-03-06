package java

import (
	"strings"
)

// GetJavaTestBuildCommand is a function for building up a javac command that can be used for building
// a java test project. This command is built up using the project configuration and the command line
// arguments passed in.
func GetJavaTestBuildCommand(configuration JavaProject, deprecation bool) JavacCommand {
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

	command.Deprecation = deprecation

	command.SourceFiles = ExtractTestSourceFileList(configuration,
		configuration.TestProject.SourceDirectory)

	if configuration.SourceVersion != "" {
		command.SourceVersion = configuration.SourceVersion
	}

	command.DebuggingInformation = ExtractDebuggingInformation(configuration)

	command.Target = configuration.Target

	return command
}

// ExtractSourceFileList is a function that reads all of the test project source files to be
// compiled from the configuration file and returns a slice of source files to be
// compiled using the javac command. Each source file has had the base path prepended
// to it when returned from the function.
func ExtractTestSourceFileList(configuration JavaProject,
	sourceDirectory string) []string {
	fileCount := len(configuration.TestProject.SourceFiles)
	fileList := make([]string, fileCount)

	for i := 0; i < fileCount; i++ {
		fileList[i] = sourceDirectory + configuration.TestProject.SourceFiles[i]
	}

	return fileList
}

// ExtractDebuggingInformation is a function that builds up the debugging information
// flag for the compiler to determine what debugging information needs to generated
// with the compiled classes.
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
