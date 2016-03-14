package project

import (
	"fmt"

	"github.com/rnowley/SonicScrewDriver/project/java"
)

// GetDocumentationBuilder is a factory function that returns an object that implements the
// DocumentationBuilder interface. The type of this DocumentationBuilder is determined by the
// configuration file that is passed in.
func GetDocumentationBuilder(configuration []byte, mode string, arguments Arguments) (DocumentationBuilder, error) {
	projectLanguage := GetProjectLanguage(configuration)
	var docBuilder DocumentationBuilder

	switch projectLanguage {
	case "java":
		docBuilder, _ = getJavaDocumentationBuilder(configuration, arguments)
		return docBuilder, nil
	default:
		return docBuilder, fmt.Errorf("GetDocumentationBuilder: the %s language is not supported", projectLanguage)
	}

}

// getJavaDocumentationBuilder is called by the factory function to return a
// documentation builder for creating javadoc documentation.
func getJavaDocumentationBuilder(configuration []byte, arguments Arguments) (DocumentationBuilder, error) {
	var proj java.JavaProject
	var projectDocumenter java.JavadocBuilder
	proj = UnmarshalJavaProject(configuration)

	command := java.GetJavadocBuildCommand(proj, arguments.Verbose)

	projectDocumenter = java.NewJavadocBuilder(command, proj)

	return projectDocumenter, nil
}
