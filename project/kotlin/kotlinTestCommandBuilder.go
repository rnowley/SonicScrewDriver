package kotlin

import "fmt"

// GetKotlincTestBuildCommand is a function for building up a kotlinc
// command that can be used for building a Kotlin test project.
// This command is built up using the project configuration that is passed in
// as a parmater to this function.
func GetKotlincTestBuildCommand(configuration KotlinProject) KotlincCommand {
	command := NewDefaultKotlincCommand()

	command.BuildTarget = configuration.BuildTarget

	if configuration.TestProject.DestinationDirectory != "" {
		command.DestinationDirectory =
			configuration.TestProject.DestinationDirectory
	}
	
	if configuration.TestProject.SourceDirectory != "" {
		command.SourceDirectory = configuration.TestProject.SourceDirectory
	}

	if configuration.TestProject.OutputFilename == "" {
		command.OutputFilename = "outTest.jar"
	} else {
		command.OutputFilename = configuration.TestProject.OutputFilename
	}

	classPathCount := len(configuration.ClassPath)

	// Add the current directory to the class path.
	command.ClassPath = append(command.ClassPath, ".")

	// Add the project's class to the command's class path.
	for i := 0; i < classPathCount; i++ {
		command.ClassPath =
			append(command.ClassPath, configuration.ClassPath[i])
	}

	classPathCount = len(configuration.TestProject.ClassPath)

	// Add the test project's class path to the command's class path
	for i := 0; i < classPathCount; i++ {
		command.ClassPath =
			append(command.ClassPath, configuration.TestProject.ClassPath[i])
	}

	command.SourceFiles = ExtractTestSourceFileList(configuration,
		command.SourceDirectory)

	return command
}

// ExtractTestSourceFileList is a function that reads all of the test project
// source files to be compiled from the configuration and returns a slice
// of source files to be compiled using the kotlinc command. Each source
// file has had the base path prepended to it when returned from the
// function.
func ExtractTestSourceFileList(configuration KotlinProject, sourceDirectory string) []string {
	fileCount := len(configuration.TestProject.SourceFiles)
	fileList := make([]string, fileCount)

	for i := 0; i < fileCount; i++ {
		fileList[i] = fmt.Sprintf("%s%s", sourceDirectory,
			configuration.TestProject.SourceFiles[i])
	}

	return fileList
}
