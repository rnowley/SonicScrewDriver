package java_tests

import (
	"testing"

	"github.com/rnowley/SonicScrewDriver/project/java"
)

func TestNewDefaultJavacCommand(t *testing.T) {
	command := java.NewDefaultJavacCommand()

	if command.CommandName != "javac" {
		t.Error(
			"For", "command.CommandName",
			"expected", "javac", "got",
			command.CommandName,
		)
	}

	if command.SourceDirectory != "./src/" {
		t.Error(
			"For", "command.SourceDirectory",
			"expected", "./src/", "got",
			command.SourceDirectory,
		)
	}

	if command.DestinationDirectory != "./build/" {
		t.Error(
			"For", "command.DestinationDirectory",
			"expected", "./build/", "got",
			command.DestinationDirectory,
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
