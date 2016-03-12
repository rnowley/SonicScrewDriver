package java

import (
	"strings"
	"testing"
)

func TestNewDefaultJavadocCommand(t *testing.T) {
	// Arrange

	// Act

	command := NewDefaultJavadocCommand()

	// Assert

	const expectedCommandName = "javadoc"
	const expectedDestinationDirectory = "./doc/"

	if command.CommandName != expectedCommandName {
		t.Error(
			"For", "TestNewDefaultJavadocCommand",
			"expected", expectedCommandName, "got",
			command.CommandName,
		)
	}

	if command.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "TestNewDefaultJavadocCommand",
			"expected", expectedDestinationDirectory, "got",
			command.DestinationDirectory,
		)
	}
}

func TestJavadocGetCommandName(t *testing.T) {

	// Arrange

	command := NewDefaultJavadocCommand()

	// Act

	commandName := command.GetCommandName()

	// Assert
	const expectedCommandName = "javadoc"

	if commandName != "javadoc" {
		t.Error(
			"For", "TestGetCommandName",
			"expected", expectedCommandName, "got",
			commandName,
		)
	}

}

func TestJavadocGetDestinationDirectory(t *testing.T) {

	// Arrange

	command := NewDefaultJavadocCommand()

	// Act

	actualDestinationDirectory := command.GetDestinationDirectory()

	// Assert
	const expectedDestinationDirectory = "./doc/"

	if actualDestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "TestJavadocGetDestinationDirectory",
			"expected", expectedDestinationDirectory, "got",
			actualDestinationDirectory,
		)
	}

}

func TestJavadocGenerateArgumentListForInstanceWithAllFieldsSet(t *testing.T) {
	const expected = "-d ./doctest/ -sourcepath " +
		"./src/a;./src/b -classpath ./lib/y;./lib/x " +
		"-linksource -private -windowtitle \"Test Title\" " +
		"-Xdoclint:syntax,accessibility -verbose " +
		"-doctitle \"Test Title\" -header \"Test header\" -bottom \"Test bottom text\""

	command := NewDefaultJavadocCommand()
	command.DestinationDirectory = "./doctest/"
	command.SourcePath = []string{"./src/a", "./src/b"}
	command.ClassPath = []string{"./lib/y", "./lib/x"}
	command.LinkSource = true
	command.AccessLevel = "private"
	command.WindowTitle = "Test Title"
	command.LintWarnings = []string {"syntax", "accessibility"}
	command.Verbose = true
	command.DocTitle = "Test Title"
	command.Header = "Test header"
	command.Bottom = "Test bottom text"

	argumentList := command.GenerateArgumentList()
	argumentString := strings.Join(argumentList, " ")

	if argumentString != expected {
		t.Error(
			"For", "TestJavadocGenerateArgumentListForInstanceWithAllFieldsSet",
			"expected", expected, "got",
			argumentString,
		)
	}
}
