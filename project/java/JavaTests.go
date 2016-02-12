package java

// JavaTests provides the structure of values required
// for a test project.
type JavaTests struct {
	// The source files that belong to the test project.
	SourceFiles []string
	// The name of the test framework to be used.
	TestFrameWork string
	// The base source directory for the test project.
	SourceDirectory string
	// ClassPath holds a list of items that are part of the class path.
	ClassPath []string
	// DestinationDirectory contains the directory that build artifacts are going to be placed in.
	DestinationDirectory string
	// The jar file to be executed
	JarFile string
	// The main class to executed.
	MainClass string
	// The arguments to be passed to the jar file or class file to
	// be executed.
	RunArguments []string
}
