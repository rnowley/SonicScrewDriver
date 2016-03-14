package java

import (
	"strings"
)

// GetJavadocRunCommand is a function for building up a javadoc command that can be used for building
// a java project. This command is built up using the project configuration and the command line
// arguments passed in.
func GetJavadocBuildCommand(configuration JavaProject, verbose bool) JavadocCommand {
	command := NewDefaultJavadocCommand()
	command.DestinationDirectory = configuration.DocumentationProject.DestinationDirectory
	command.Verbose = verbose
	command.SourcePath = ExtractJavadocSourceFileList(configuration)
	command.ClassPath = strings.Join(configuration.DocumentationProject.ClassPath, ";")
	command.LinkSource = configuration.DocumentationProject.LinkSource
	command.WindowTitle = configuration.DocumentationProject.WindowTitle
	command.DocTitle = configuration.DocumentationProject.DocTitle
	command.Header = configuration.DocumentationProject.Header
	command.Bottom = configuration.DocumentationProject.Bottom

	return command
}

// ExtractSourceFileList is a function that reads all of the source files to be
// compiled from the configuration file and returns a slice of source files to be
// compiled using the javac command. Each source file has had the base path appended
// to it when returned from the function.
func ExtractJavadocSourceFileList(configuration JavaProject) []string {
	fileCount := len(configuration.DocumentationProject.SourcePath)
	fileList := make([]string, fileCount)

	for i := 0; i < fileCount; i++ {
		fileList[i] = configuration.DocumentationProject.SourcePath[i]
	}

	return fileList
}
