package kotlin

import (
	"strings"
	"testing"
)

func TestNewDefaultKotlinCommand(t *testing.T) {
	// Arrange

	// Act
	commandToTest := NewDefaultKotlinCommand()

	// Assert

	const expectedCommandName = "java"

	if commandToTest.CommandName != expectedCommandName {
		t.Error(
			"For", "TestNewDefaultKotlinCommand",
			"expected", expectedCommandName, "got",
			commandToTest.CommandName,
		)
	}

	const expectedJarFile = "./build/out.jar"

	if commandToTest.JarFile != expectedJarFile {
		t.Error(
			"For", "TestNewDefaultKotlinCommand",
			"expected", expectedJarFile, "got",
			commandToTest.JarFile,
		)
	}

	if commandToTest.ClassPath == nil {
		t.Error(
			"For", "TestNewDefaultKotlinCommand",
			"expected", "!= nil", "got",
			commandToTest.ClassPath,
		)
	}
}

func TestKotlinGetCommandName(t *testing.T) {
	// Arrange

	// Act
	commandToTest := NewDefaultKotlinCommand()
	actualCommandName := commandToTest.GetCommandName()

	// Assert
	const expectedCommandName = "java"

	if actualCommandName != expectedCommandName {
		t.Error(
			"For", "TestKotlincGetCommandName",
			"expected", expectedCommandName, "got",
			actualCommandName,
		)
	}

}

func TestKotlinGenerateArgumentListForDefaultInstance(t *testing.T) {
	// Arrange

	// Act
	commandToTest := NewDefaultKotlinCommand()

	argumentList := commandToTest.GenerateArgumentList()
	actualArgumentString := strings.Join(argumentList, " ")

	// Assert

	const expectedArgumentString = "-jar ./build/out.jar"

	if actualArgumentString != expectedArgumentString {
		t.Error(
			"For", "TestKotlinGenerateArgumentListForDefaultInstance",
			"expected", expectedArgumentString, "got",
			actualArgumentString,
		)
	}
}

func TestKotlinGenerateArgumentListForInstanceWithAllFieldsSetJarFile(t *testing.T) {
	// Arrange

	command := NewDefaultKotlinCommand()
	command.ClassPath = []string{".", "./lib/a.jar", "./lib/x/b.jar"}
	command.JarFile = "testout.jar"
	command.RunArguments = []string{"arg1", "arg2"}

	// Act

	argumentList := command.GenerateArgumentList()
	actualArgumentString := strings.Join(argumentList, " ")

	// Assert

	const expectedArgumentString = "-cp .:./lib/a.jar:./lib/x/b.jar -jar testout.jar " +
		"arg1 arg2"

	if actualArgumentString != expectedArgumentString {
		t.Error(
			"For", "TestKotlinGenerateArgumentListForInstanceWithAllFieldsSet",
			"expected", expectedArgumentString, "got",
			actualArgumentString,
		)
	}
}

func TestKotlinGenerateArgumentListForInstanceWithAllFieldsSetMainClass(t *testing.T) {
	// Arrange

	command := NewDefaultKotlinCommand()
	command.ClassPath = []string{".", "./lib/a.jar", "./lib/x/b.jar"}
	command.JarFile = ""
	command.MainClass = "testout"
	command.RunArguments = []string{"arg1", "arg2"}

	// Act

	argumentList := command.GenerateArgumentList()
	actualArgumentString := strings.Join(argumentList, " ")

	// Assert

	const expectedArgumentString = "-cp .:./lib/a.jar:./lib/x/b.jar testout " +
		"arg1 arg2"

	if actualArgumentString != expectedArgumentString {
		t.Error(
			"For", "TestKotlinGenerateArgumentListForInstanceWithAllFieldsSet",
			"expected", expectedArgumentString, "got",
			actualArgumentString,
		)
	}
}

func TestKotlinGenerateArgumentListForInstanceWithAllFieldsSetJarFileAndMainClass(t *testing.T) {
	// Arrange

	command := NewDefaultKotlinCommand()
	command.ClassPath = []string{".", "./lib/a.jar", "./lib/x/b.jar"}
	command.JarFile = "testout.jar"
	command.MainClass = "testout"
	command.RunArguments = []string{"arg1", "arg2"}

	// Act

	argumentList := command.GenerateArgumentList()
	actualArgumentString := strings.Join(argumentList, " ")

	// Assert

	const expectedArgumentString = "-cp .:./lib/a.jar:./lib/x/b.jar -jar testout.jar " +
		"arg1 arg2"

	if actualArgumentString != expectedArgumentString {
		t.Error(
			"For", "TestKotlinGenerateArgumentListForInstanceWithAllFieldsSet",
			"expected", expectedArgumentString, "got",
			actualArgumentString,
		)
	}
}
