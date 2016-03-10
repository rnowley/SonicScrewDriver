package java

// JavaDocumentation provides the structure of values required
// for generating JavaDoc documentation.
// Author is used to determine if the documentation should contain an author
// section.
// ClassPath specifies the paths where javadoc searches for referenced classes.
// DestinationDirectory specifies where to save the generated HTML files.
// AccessLevel indicates the access level of the members to be displayed in the
// documentation output (can be public, protected, package, private).
// SourcePath specifies the search paths for finding source files.
// LinkSource when set to true, creates a HTML version of each source file
// (with line numbers) which is linked from the HTML documentation.
type JavaDocumentation struct {
	Author               bool
	ClassPath            []string
	DestinationDirectory string
	AccessLevel          string
	SourcePath           []string
	LinkSource           bool
}
