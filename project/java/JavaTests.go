package java

// JavaTests provides the structure of values required
// for a test project.
// SourceFiles holds the source files that belong to the test project.
// TestFrameWork holds the name of the test framework to be used.
// The base source directory for the test project.
// ClassPath holds a list of items that are part of the class path.
// DestinationDirectory contains the directory that build artifacts are going to be placed in.
// The jar file to be executed
// The main class to executed.
// The arguments to be passed to the jar file or class file to
// be executed.
type JavaTests struct {
	SourceFiles          []string
	TestFrameWork        string
	SourceDirectory      string
	ClassPath            []string
	DestinationDirectory string
	JarFile              string
	MainClass            string
	RunArguments         []string
}
