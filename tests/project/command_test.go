package project_test

import (
	"github.com/rnowley/sonicScrewDriver/project"
	"testing"
)

func TestNewDefaultCommand(t *testing.T) {
	c := project.NewDefaultCommand()
	//d := *c

	if c.CommandName != "javac" {
		t.Error(
			"For", "command.Name",
			"expected", "javac",
			"got", c.CommandName,
		)
	}
}
