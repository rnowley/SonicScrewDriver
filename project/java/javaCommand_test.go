package java

import (
	"strings"
	"testing"
)

func TestNewDefaultJavaCommand(t *testing.T) {
	const expected = "java"

	command := NewDefaultJavaCommand()

	if command.CommandName != expected {
		t.Error(
			"For", "command.CommandName",
			"expected", expected, "got",
			command.CommandName,
		)
	}

	if command.ClassPath == nil {
		t.Error(
			"For", "command.ClassPath",
			"expected", "!= nil", "got",
			command.ClassPath,
		)
	}

}

func TestJavaGetCommandName(t *testing.T) {
	const expected = "java"

	command := NewDefaultJavaCommand()
	actual := command.GetCommandName()

	if actual != expected {
		t.Error(
			"For", "command.GetCommandName()",
			"expected", expected, "got",
			actual,
		)
	}

}

func TestJavaGenerateArgumentListForDefaultInstance(t *testing.T) {
	const expected = ""

	command := NewDefaultJavaCommand()

	argumentList := command.GenerateArgumentList()
	actual := strings.Join(argumentList, " ")

	if actual != expected {
		t.Error(
			"For", "command.GetArgumentList()",
			"expected", expected, "got",
			actual,
		)
	}

}

func TestJavaGenerateArgumentListForInstanceWithAllFieldsSetExceptMainClass(t *testing.T) {
	const expected = "-cp ./lib/a.jar:./lib/b.jar -jar c.jar " +
		"arg1 arg2"

	command := NewDefaultJavaCommand()
	command.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	command.JarFile = "c.jar"
	command.RunArguments = []string{"arg1", "arg2"}

	argumentList := command.GenerateArgumentList()
	actual := strings.Join(argumentList, " ")

	if actual != expected {
		t.Error(
			"For", "command.GetArgumentList()",
			"expected", expected, "got",
			actual,
		)
	}
}

func TestJavaGenerateArgumentListForInstanceWithAllFieldsSetExceptJarFile(t *testing.T) {
	const expected = "-cp ./lib/a.jar:./lib/b.jar Main.Class " +
		"arg1 arg2"

	command := NewDefaultJavaCommand()
	command.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	command.MainClass = "Main.Class"
	command.RunArguments = []string{"arg1", "arg2"}

	argumentList := command.GenerateArgumentList()
	actual := strings.Join(argumentList, " ")

	if actual != expected {
		t.Error(
			"For", "command.GetArgumentList()",
			"expected", expected, "got",
			actual,
		)
	}
}

func TestJavaGenerateArgumentListForInstanceWithAllFieldsSet(t *testing.T) {
	const expected = "-cp ./lib/a.jar:./lib/b.jar -jar c.jar " +
		"arg1 arg2"

	command := NewDefaultJavaCommand()
	command.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	command.JarFile = "c.jar"
	command.MainClass = "Main.Class"
	command.RunArguments = []string{"arg1", "arg2"}

	argumentList := command.GenerateArgumentList()
	actual := strings.Join(argumentList, " ")

	if actual != expected {
		t.Error(
			"For", "command.GetArgumentList()",
			"expected", expected, "got",
			actual,
		)
	}
}
