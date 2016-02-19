package kotlin

// KotlinProject is a struct for holding the information required for
// building or running a Kotlin Project.
type KotlinProject struct {
	Name            string
	Version         string
	Description     string
	Language        string
	Destination     string
	SourceDirectory string
	KotlinHome      string
	ClassPath       []string
	SourceFiles     []string
}
