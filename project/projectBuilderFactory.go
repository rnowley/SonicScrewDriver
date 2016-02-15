package project

import (
	"encoding/json"
	"fmt"

	"github.com/rnowley/SonicScrewDriver/project/csharp"
	"github.com/rnowley/SonicScrewDriver/project/java"
)

// GetProjectBuilder is a factory function that returns an object that implements the
// ProjectBuilder interface. The type of this ProjectBuilder is determined by the
// configuration file that is passed in.
func GetProjectBuilder(configurationFile []byte, mode string, arguments Arguments) (ProjectBuilder, error) {
	projectLanguage := GetProjectLanguage(configurationFile)
	var projectBuilder ProjectBuilder

	switch projectLanguage {
	case "csharp":
		projectBuilder, _ = getCSharpProjectBuilder(configurationFile)
		return projectBuilder, nil
	case "java":
		projectBuilder, _ = getJavaProjectBuilder(configurationFile, mode, arguments)
		return projectBuilder, nil
	default:
		return projectBuilder, fmt.Errorf("GetProjectBuilder: the %s language is not supported", projectLanguage)
	}

}

// GetProjectLanguage is a function for retrieving the value that
// determines the programming language that the project is written
// in.
/*func GetProjectLanguage(file []byte) string {
	var projectLanguage ProjectLanguage

	if err := json.Unmarshal(file, &projectLanguage); err != nil {
		panic(err)
	}

	return projectLanguage.Language
}*/

func buildCSharpCommand(proj csharp.CSharpProject) csharp.CSharpCommand {
	command := csharp.BuildCommand(proj)
	return command
}

func getCSharpProjectBuilder(configurationFile []byte) (ProjectBuilder, error) {
	var proj csharp.CSharpProject
	var projectBuilder csharp.CSharpProjectBuilder

	proj, err := unmarshalCSharpProject(configurationFile)

	if err != nil {
		return projectBuilder, err
	}

	command := buildCSharpCommand(proj)
	projectBuilder = csharp.New(command, proj)

	return projectBuilder, nil
}

func getJavaProjectBuilder(configurationFile []byte, mode string, arguments Arguments) (ProjectBuilder, error) {
	var proj java.JavaProject
	var projectBuilder java.JavaProjectBuilder

	proj = unmarshalJavaProject(configurationFile)

	var command java.JavacCommand

	switch mode {
	case "build":
		command = java.GetJavaBuildCommand(proj, arguments.Deprecation)
	case "build-test":
		command = java.GetJavaTestBuildCommand(proj, arguments.Deprecation)
	default:
		return projectBuilder, fmt.Errorf("getJavaProjectBuilder: the %s 'mode' is not supported", mode)
	}

	projectBuilder = java.NewProjectBuilder(command, proj)

	return projectBuilder, nil
}

// UnmarshalCSharpProject is a function that takes in the JSON representation of
// a CSharp project and transforms this into a CSharpProject object.
func unmarshalCSharpProject(projectFile []byte) (csharp.CSharpProject, error) {
	var proj csharp.CSharpProject

	if err := json.Unmarshal(projectFile, &proj); err != nil {
		return proj, err
	}

	return proj, nil
}

// UnmarshalJavaProject is a function that takes in the JSON representation of
// a Java project and transforms this into a JavaProject object.
func unmarshalJavaProject(projectFile []byte) java.JavaProject {
	var proj java.JavaProject

	if err := json.Unmarshal(projectFile, &proj); err != nil {
		panic(err)
	}
	fmt.Println(proj)
	return proj
}
