package project

import "time"

// BuildStatistics is a struct for holding the exit code and the elapsed time
// of a build operation.
type BuildStatistics struct {
	ExitCode    int
	ElapsedTime time.Time
}
