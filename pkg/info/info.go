package info

import "fmt"

var (
	// set these during go build/run via -ldflags
	version = "unknown"
	commit  = "unknown"
	branch  = "unknown"
)

func Info() string {
	return fmt.Sprintf("version %s (branch: %s, revision: %s)", version, branch, commit)
}

func Version() string {
	return version
}
