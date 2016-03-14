package project

type DocumentationBuilder interface {

	// BuildDocumentation is where you implement actions related to building
	// project documentation.
	BuildDocumentation(verbose bool) error
}
