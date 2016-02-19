package project

import (
	"encoding/json"
	"fmt"

	"github.com/rnowley/SonicScrewDriver/project/csharp"
	"github.com/rnowley/SonicScrewDriver/project/java"
	"github.com/rnowley/SonicScrewDriver/project/kotlin"
)

// GetProjectLanguage is a function for retrieving the value that
// determines the programming language that the project is written
// in.
func GetProjectLanguage(file []byte) string {
	var projectLanguage ProjectLanguage

	if err := json.Unmarshal(file, &projectLanguage); err != nil {
		panic(err)
	}

	return projectLanguage.Language
}

// UnmarshalCSharpProject is a function that takes in the JSON representation of
// a CSharp project and transforms this into a CSharpProject object.
func UnmarshalCSharpProject(projectFile []byte) (csharp.CSharpProject, error) {
	var proj csharp.CSharpProject

	if err := json.Unmarshal(projectFile, &proj); err != nil {
		return proj, err
	}

	return proj, nil
}

// UnmarshalJavaProject is a function that takes in the JSON representation of
// a Java project and transforms this into a JavaProject object.
func UnmarshalJavaProject(projectFile []byte) java.JavaProject {
	var proj java.JavaProject

	if err := json.Unmarshal(projectFile, &proj); err != nil {
		panic(err)
	}

	fmt.Println(proj)
	return proj
}

// UnmarshalKotlinProject is a function that takes in the JSON representation of
// a Kotlin project and transforms this into a KotlinProject object.
func UnmarshalKotlinProject(projectFile []byte) kotlin.KotlinProject {
	var proj kotlin.KotlinProject

	if err := json.Unmarshal(projectFile, &proj); err != nil {
		panic(err)
	}

	fmt.Println(proj)
	return proj
}
