package java

import (
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
