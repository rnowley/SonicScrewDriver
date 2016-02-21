package project

import (
	"fmt"

	"github.com/rnowley/SonicScrewDriver/project/csharp"
	"github.com/rnowley/SonicScrewDriver/project/java"
	"github.com/rnowley/SonicScrewDriver/project/kotlin"
)

// GetProjectRunner is a factory function that returns an object that implements the
// ProjectRunner interface. The type of this ProjectRunner is determined by the
// configuration file that is passed in.
func GetProjectRunner(configurationFile []byte, mode string, arguments Arguments) (ProjectRunner, error) {
	projectLanguage := GetProjectLanguage(configurationFile)
	var projectRunner ProjectRunner

	switch projectLanguage {
	case "csharp":
		projectRunner, _ = getCSharpProjectRunner(configurationFile, mode, arguments)
		return projectRunner, nil
	case "java":
		projectRunner, _ = getJavaProjectRunner(configurationFile, mode, arguments)
		return projectRunner, nil
	case "kotlin":
		projectRunner, _ = getKotlinProjectRunner(configurationFile, mode, arguments)
		return projectRunner, nil
	default:
		return projectRunner, fmt.Errorf("GetProjectRunner: the %s language is not supported", projectLanguage)
	}

}

// getCSharpProjectRunner retrieves a CSharpProjectRunner that is configured to the
// specifications passed in in the configuration file, mode and arguments.
func getCSharpProjectRunner(configurationFile []byte, mode string, arguments Arguments) (ProjectRunner, error) {
	var proj csharp.CSharpProject
	var projectRunner csharp.CSharpProjectRunner

	proj, _ = UnmarshalCSharpProject(configurationFile)

	var command csharp.MonoCommand

	switch mode {
	case "run":
		command = csharp.GetCSharpRunCommand(proj)
	case "run-tests":
		command = csharp.GetCSharpRunTestCommand(proj)
	default:
		return projectRunner, fmt.Errorf("getCSharpProjectRunner: the %s 'mode' is not supported", mode)
	}

	projectRunner = csharp.NewProjectRunner(command, proj)

	return projectRunner, nil
}

// getJavaProjectRunner retrieves a JavaProjectRunner that is configured to the
// specifications passed in in the configuration file, mode and arguments.
func getJavaProjectRunner(configurationFile []byte, mode string, arguments Arguments) (ProjectRunner, error) {
	var proj java.JavaProject
	var projectRunner java.JavaProjectRunner

	proj = UnmarshalJavaProject(configurationFile)

	var command java.JavaCommand

	switch mode {
	case "run":
		command = java.GetJavaRunCommand(proj)
	case "run-tests":
		command = java.GetJavaRunTestCommand(proj)
	default:
		return projectRunner, fmt.Errorf("getJavaProjectRunner: the %s 'mode' is not supported", mode)
	}

	projectRunner = java.NewProjectRunner(command, proj)

	return projectRunner, nil
}

func getKotlinProjectRunner(configurationFile []byte, mode string, arguments Arguments) (ProjectRunner, error) {
	var proj kotlin.KotlinProject
	var projectRunner kotlin.KotlinProjectRunner
	var command kotlin.KotlinCommand

	proj = UnmarshalKotlinProject(configurationFile)

	switch mode {
	case "run":
		command = kotlin.GetKotlinRunCommand(proj)
	case "run-tests":
		return projectRunner, fmt.Errorf("getKotlinProjectRunner: the %s mode has not been implemented", mode)
	default:
		return projectRunner, fmt.Errorf("getKotlinProjectRunner: the %s 'mode' is not supported", mode)
	}

	projectRunner = kotlin.NewKotlinProjectRunner(command, proj)

	return projectRunner, nil

}
