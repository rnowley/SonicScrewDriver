package project

// ProjectBuilder provides an interface for objects concerned with
// building a particular type of project.
type ProjectBuilder interface {

	// ExecutePreBuildTasks is where you implement actions that need to take place
	// before the build starts.
	ExecutePreBuildTasks(verbose bool) error

	// BuildProject is where you implement actions related to the build itself.
	BuildProject(verbose bool) error

	// ExecutePostBuildTasks is where you implement actions that need to take place
	// after the build has completed.
	ExecutePostBuildTasks(verbose bool) error
}
