package java

// JavaProject is a struct for holding the information required for
// building or running a Java project.
type JavaProject struct {
	// Name is the name of the project.
	Name string

	// Version is used for holding the project version number.
	Version string

	// Description holds a description of the project.
	Description string

	// Language holds the name of the programming language that the project is written in
	Language string

	// DestinationDirectory contains the directory that build artifacts are going to be placed in.
	DestinationDirectory string

	// ClassPath holds a list of items that are part of the class path.
	ClassPath []string

	// SourceFiles holds a list of all of the files to be compiled into the project.
	SourceFiles []string

	// SourceVersion holds the value of the source code version accepted by the compiler.
	SourceVersion string

	// DebuggingInformation holds information about the debugging information to be generated.
	DebuggingInformation []string

	// The jar file to be executed
	JarFile string

	// The main class to executed.
	MainClass string

	// The arguments to be passed to the jar file or class file to
	// be executed.
	RunArguments []string

	// The lint warnings to be enabled for compilation.
	LintWarnings []string

	// Sets the source file encoding.
	Encoding string

	// TestProject holds information about the tests for this project.
	TestProject JavaTests
}
