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
	command.SourcePath = strings.Join(configuration.DocumentationProject.SourcePath, ";")
	command.ClassPath = strings.Join(configuration.DocumentationProject.ClassPath, ";")
	command.LinkSource = configuration.DocumentationProject.LinkSource
	command.WindowTitle = configuration.DocumentationProject.WindowTitle
	command.DocTitle = configuration.DocumentationProject.DocTitle
	command.Header = configuration.DocumentationProject.Header
	command.Bottom = configuration.DocumentationProject.Bottom

	return command
}
