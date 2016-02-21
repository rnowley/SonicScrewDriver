package kotlin

// KotlinTests provides the information to
// build a test project.
type KotlinTests struct {
	SourceFiles          []string
	SourceDirectory      string
	DestinationDirectory string
	OutputFilename       string
	ClassPath            []string
	TestRunner           string
	RunArguments         []string
}
