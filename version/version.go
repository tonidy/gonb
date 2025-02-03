package version

import "github.com/janpfeifer/gonb/internal/version"

// AppVersion contains version and Git commit information.
var AppVersion = version.AppVersion(GitTag, GitCommitHash)
