package project

import (
	"fmt"

	"github.com/rnowley/SonicScrewDriver/project/java"
)

// GetProjectRunner is a factory function that returns an object that implements the
// ProjectRunner interface. The type of this ProjectRunner is determined by the
// configuration file that is passed in.
func GetProjectRunner(configurationFile []byte, mode string, arguments Arguments) (ProjectRunner, error) {
	projectLanguage := GetProjectLanguage(configurationFile)
	var projectRunner ProjectRunner

	switch projectLanguage {
	case "java":
		projectRunner, _ = getJavaProjectRunner(configurationFile, mode, arguments)
		return projectRunner, nil
	default:
		return projectRunner, fmt.Errorf("GetProjectRunner: the %s language is not supported", projectLanguage)
	}

}

// getJavaProjectRunner retrieves a JavaProjectRunner the is configured to the
// specifications passed in in the configuration file, mode and arguments.
func getJavaProjectRunner(configurationFile []byte, mode string, arguments Arguments) (ProjectRunner, error) {
	var proj java.JavaProject
	var projectRunner java.JavaProjectRunner

	proj = UnmarshalJavaProject(configurationFile)

	var command java.JavaCommand

	switch mode {
	case "run-tests":
		command = java.GetJavaRunTestCommand(proj)
		fmt.Printf("Command: %s", command)
	default:
		return projectRunner, fmt.Errorf("getJavaProjectRunner: the %s 'mode' is not supported", mode)
	}

	projectRunner = java.NewProjectRunner(command, proj)

	return projectRunner, nil
}
