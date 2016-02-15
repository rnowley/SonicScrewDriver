package project

type ProjectRunner interface {

	// RunProject is where you implement actions related to running a
	// project.
	RunProject() error
}
