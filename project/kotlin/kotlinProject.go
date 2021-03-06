package kotlin

import "github.com/rnowley/SonicScrewDriver/project/common"

// KotlinProject is a struct for holding the information required for
// building or running a Kotlin Project.
type KotlinProject struct {
	Name                 string
	Version              string
	Description          string
	Language             string
	DestinationDirectory string
	JarFile              string
	SourceDirectory      string
	OutputFilename       string
	KotlinHome           string
	ClassPath            []string
	SourceFiles          []string
	Resources            []common.Resource
	BuildTarget          string
	RunArguments         []string
	TestProject          KotlinTests
}
