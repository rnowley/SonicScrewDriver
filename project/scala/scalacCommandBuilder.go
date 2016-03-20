package scala

// GetScalacBuildCommand is a function for building up a sclac command that can be used for building
// a scala project. This command is built up using the project configuration and the command line
// arguments passed in.
func GetScalacBuildCommand(configuration ScalaProject, deprecation bool,
	verbose bool) ScalacCommand {
	command := NewDefaultScalacCommand()

	if len(configuration.ClassPath) != 0 {
		command.ClassPath = configuration.ClassPath
	}

	command.DebuggingInformation = configuration.DebuggingInformation
	command.Deprecation = deprecation

	if configuration.DestinationDirectory != "" {
		command.DestinationDirectory = configuration.DestinationDirectory
	}

	command.Encoding = configuration.Encoding
	command.NoWarnings = configuration.NoWarnings
	command.Optimise = configuration.Optimise

	if configuration.SourceDirectory != "" {
		command.SourceDirectory = configuration.SourceDirectory
	}

	command.SourceFiles = ExtractSourceFileList(configuration, command.SourceDirectory)
	command.Target = configuration.Target
	command.Verbose = verbose

	return command
}

// ExtractSourceFileList is a function that reads all of the source files to be
// compiled from the configuration file and returns a slice of source files to be
// compiled using the scalac command. Each source file has had the base path appended
// to it when returned from the function.
func ExtractSourceFileList(configuration ScalaProject,
	sourceDirectory string) []string {
	fileCount := len(configuration.SourceFiles)
	fileList := make([]string, fileCount)

	for i := 0; i < fileCount; i++ {
		fileList[i] = sourceDirectory + configuration.SourceFiles[i]
	}

	return fileList
}
