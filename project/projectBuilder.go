package project

type ProjectBuilder interface {
    ExecutePreBuildTasks() error
    BuildProject() error
    ExecutePostBuildTasks() error
}