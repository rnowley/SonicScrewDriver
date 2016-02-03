package project

type ProjectBuilder interface {
    BuildProject(command Command) BuildStatistics
    PostBuildTasks(command Command)
}