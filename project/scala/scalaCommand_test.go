package scala

import (
	"strings"
	"testing"
)

func TestNewDefaultScalaCommand(t *testing.T) {
	// Arrange

	// Act
	command := NewDefaultScalaCommand()

	// Assert
	const expectedCommandName = "scala"

	if command.CommandName != expectedCommandName {
		t.Error(
			"For", "TestNewDefaultScalaCommand",
			"expected", expectedCommandName, "got",
			command.CommandName,
		)
	}

	if command.ClassPath == nil {
		t.Error(
			"For", "TestNewDefaultScalaCommand",
			"expected", "!= nil", "got",
			command.ClassPath,
		)
	}
}

func TestScalaGetCommandName(t *testing.T) {
	// Arrange
	command := NewDefaultScalaCommand()

	// Act
	actualCommandName := command.GetCommandName()

	// Assert
	const expectedCommandName = "scala"

	if actualCommandName != expectedCommandName {
		t.Error(
			"For", "TestScalaGetCommandName",
			"expected", expectedCommandName, "got",
			actualCommandName,
		)
	}

}

func TestScalaGenerateArgumentListForDefaultInstance(t *testing.T) {
	// Arrange
	command := NewDefaultScalaCommand()

	// Act
	argumentList := command.GenerateArgumentList()

	// Assert
	const expectedArgumentList = ""

	actualArgumentList := strings.Join(argumentList, " ")

	if actualArgumentList != expectedArgumentList {
		t.Error(
			"For", "TestScalaGenerateArgumentListForDefaultInstance",
			"expected", expectedArgumentList, "got",
			actualArgumentList,
		)
	}

}

func TestScalaGenerateArgumentListForInstanceWithAllFieldsSetExceptMainClass(t *testing.T) {
	// Arrange
	command := NewDefaultScalaCommand()
	command.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	command.JarFile = "c.jar"
	command.RunArguments = []string{"arg1", "arg2"}

	// Act
	argumentList := command.GenerateArgumentList()
	actualArgumentList := strings.Join(argumentList, " ")

	// Assert
	const expectedArgumentList = "-classpath ./lib/a.jar:./lib/b.jar -jar c.jar " +
		"arg1 arg2"

	if actualArgumentList != expectedArgumentList {
		t.Error(
			"For", "TestScalaGenerateArgumentListForInstanceWithAllFieldsSetExceptMainClass",
			"expected", expectedArgumentList, "got",
			actualArgumentList,
		)
	}
}

func TestScalaGenerateArgumentListForInstanceWithAllFieldsSetExceptJarFile(t *testing.T) {
	// Arrange
	command := NewDefaultScalaCommand()
	command.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	command.MainClass = "Main.Class"
	command.RunArguments = []string{"arg1", "arg2"}

	// Act
	argumentList := command.GenerateArgumentList()
	actualArgumentList := strings.Join(argumentList, " ")

	// Assert
	const expectedArgumentList = "-classpath ./lib/a.jar:./lib/b.jar Main.Class " +
		"arg1 arg2"

	if actualArgumentList != expectedArgumentList {
		t.Error(
			"For", "TestScalaGenerateArgumentListForInstanceWithAllFieldsSetExceptJarFile",
			"expected", expectedArgumentList, "got",
			actualArgumentList,
		)
	}
}

func TestScalaGenerateArgumentListForInstanceWithAllFieldsSet(t *testing.T) {
	// Arrange
	command := NewDefaultScalaCommand()
	command.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	command.JarFile = "c.jar"
	command.MainClass = "Main.Class"
	command.RunArguments = []string{"arg1", "arg2"}

	// Act
	argumentList := command.GenerateArgumentList()
	actualArgumentList := strings.Join(argumentList, " ")

	// Assert
	const expectedArgumentList = "-classpath ./lib/a.jar:./lib/b.jar -jar c.jar " +
		"arg1 arg2"

	if actualArgumentList != expectedArgumentList {
		t.Error(
			"For", "TestScalaGenerateArgumentListForInstanceWithAllFieldsSet",
			"expected", expectedArgumentList, "got",
			actualArgumentList,
		)
	}
}
