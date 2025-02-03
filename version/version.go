package version

import (
	"embed"
	"strings"

	"github.com/janpfeifer/gonb/internal/version"
)

// AppVersion contains version and Git commit information.
var AppVersion = version.AppVersion(getGitVersion(), getGitCommitHash())

//go:embed *.txt
var textFiles embed.FS

func getGitVersion() string {
	content, err := textFiles.ReadFile("version.txt")
	if err != nil {
		return "0.1.0"
	}
	return strings.TrimSpace(string(content))
}

func getGitCommitHash() string {
	content, err := textFiles.ReadFile("hash.txt")
	if err != nil {
		return "default-sha1-hash"
	}
	return strings.TrimSpace(string(content))
}
