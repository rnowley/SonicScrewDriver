package java

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
