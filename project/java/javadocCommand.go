package java

import (
	"fmt"
	"strings"
)

// JavadocCommand provides information for running the
// javadoc command.
type JavadocCommand struct {
	CommandName          string
	DestinationDirectory string
	SourceDirectory      string
	SourcePath           []string
	ClassPath            string
	LinkSource           bool
	AccessLevel          string
	LintWarnings         []string
	WindowTitle          string
	Verbose              bool
	DocTitle             string
	Header               string
	Bottom               string
}

// NewDefaultJavadocCommand returns a Javadoc command with some default values.
func NewDefaultJavadocCommand() JavadocCommand {
	var command JavadocCommand
	command.CommandName = "javadoc"
	command.DestinationDirectory = "./doc/"
	command.LintWarnings = make([]string, 0)
	return command
}

// GetCommandName is a method on a JavadocCommand which accesses the name of the command
// to be run.
func (command JavadocCommand) GetCommandName() string {
	return command.CommandName
}

// GetDestination is a method which returns the the destination directory where the
// compiler's output is going to be copied to.
func (command JavadocCommand) GetDestinationDirectory() string {
	return command.DestinationDirectory
}

func (command JavadocCommand) GenerateArgumentList() []string {
	argumentArray := make([]string, 0)
	argumentArray = append(argumentArray, "-d", command.DestinationDirectory)

	if len(command.SourcePath) != 0 {
		argumentArray = append(argumentArray, command.SourcePath...)
	}

	if len(command.ClassPath) != 0 {
		argumentArray = append(argumentArray, "-classpath", command.ClassPath)
	}

	if command.LinkSource {
		argumentArray = append(argumentArray, "-linksource")
	}

	if command.AccessLevel != "" {
		argumentArray = append(argumentArray, fmt.Sprintf("-%s", command.AccessLevel))
	}

	if command.WindowTitle != "" {
		argumentArray = append(argumentArray, "-windowtitle",
			fmt.Sprintf("\"%s\"", command.WindowTitle))
	}

	if len(command.LintWarnings) != 0 {
		argumentArray = append(argumentArray, fmt.Sprintf("-Xdoclint:%s", strings.Join(command.LintWarnings, ",")))
	}

	if command.Verbose {
		argumentArray = append(argumentArray, "-verbose")
	}

	if command.DocTitle != "" {
		argumentArray = append(argumentArray, "-doctitle",
			fmt.Sprintf("\"%s\"", command.DocTitle))
	}

	if command.Header != "" {
		argumentArray = append(argumentArray, "-header",
			fmt.Sprintf("\"%s\"", command.Header))
	}

	if command.Bottom != "" {
		argumentArray = append(argumentArray, "-bottom",
			fmt.Sprintf("\"%s\"", command.Bottom))
	}

	return argumentArray

}

func (command JavadocCommand) String() string {
	arguments := strings.Join(command.GenerateArgumentList(), " ")
	return fmt.Sprintf("%s %s", command.GetCommandName(), arguments)
}
