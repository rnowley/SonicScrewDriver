package scala

func GetScalacTestBuildCommand(configuration ScalaProject, deprecation bool,
	verbose bool) ScalacCommand {
	command := NewDefaultScalacCommand()
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

	command.SourceFiles = ExtractTestSourceFileList(configuration,
		configuration.TestProject.SourceDirectory)

	command.Deprecation = deprecation
	command.Verbose = verbose

	return command
}

// ExtractSourceFileList is a function that reads all of the test project source files to be
// compiled from the configuration file and returns a slice of source files to be
// compiled using the scalac command. Each source file has had the base path prepended
// to it when returned from the function.
func ExtractTestSourceFileList(configuration ScalaProject,
	sourceDirectory string) []string {
	fileCount := len(configuration.TestProject.SourceFiles)
	fileList := make([]string, fileCount)

	for i := 0; i < fileCount; i++ {
		fileList[i] = sourceDirectory + configuration.TestProject.SourceFiles[i]
	}

	return fileList
}
