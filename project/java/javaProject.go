package java

// JavaProject is a struct for holding the information required for
// building or running a Java project.
// Name is the name of the project.
// Version is used for holding the project version number.
// Description holds a description of the project.
// Language holds the name of the programming language that the project is written in.
// The root directory for the project's source code.
// DestinationDirectory contains the directory that build artifacts are going to be placed in.
// ClassPath holds a list of items that are part of the class path.
// SourceFiles holds a list of all of the files to be compiled into the project.
// SourceVersion holds the value of the source code version accepted by the compiler.
// DebuggingInformation holds information about the debugging information to be generated.
// The jar file to be executed.
// The main class to executed.
// The arguments to be passed to the jar file or class file to
// be executed.
// The lint warnings to be enabled for compilation.
// Sets the source file encoding.
// Sets the specific version of the Java VM to target.
// TestProject holds information about the tests for this project.
// Holds information for generating documentation.
type JavaProject struct {
	Name                 string
	Version              string
	Description          string
	Language             string
	SourceDirectory      string
	DestinationDirectory string
	ClassPath            []string
	SourceFiles          []string
	SourceVersion        string
	DebuggingInformation []string
	JarFile              string
	MainClass            string
	RunArguments         []string
	LintWarnings         []string
	Encoding             string
	Target               string
	TestProject          JavaTests
	DocumentationProject JavaDocumentation
}
