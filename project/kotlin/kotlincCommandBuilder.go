package kotlin

// GetKotlncBuildCommand is a function for building up a kotlinc command that can be used for building
// a kotlin project. This command is built up using the project configuration and the command line
// arguments passed in.
func GetKotlincBuildCommand(configuration KotlinProject) KotlincCommand {
	command := NewDefaultKotlincCommand()

	if len(configuration.ClassPath) != 0 {
		command.ClassPath = configuration.ClassPath
	}

	if configuration.Destination != "" {
		command.DestinationDirectory = configuration.Destination
	}

	command.SourceFiles = ExtractSourceFileList(configuration, command.SourceDirectory)

	return command
}

// ExtractSourceFileList is a function that reads all of the source files to be
// compiled from the configuration file and returns a slice of source files to be
// compiled using the kotlinc command. Each source file has had the base path appended
// to it when returned from the function.
func ExtractSourceFileList(configuration KotlinProject,
	sourceDirectory string) []string {
	fileCount := len(configuration.SourceFiles)
	fileList := make([]string, fileCount)

	for i := 0; i < fileCount; i++ {
		fileList[i] = sourceDirectory + configuration.SourceFiles[i]
	}

	return fileList
}
