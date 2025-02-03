package main

import (
	"os"

	"github.com/janpfeifer/gonb/gonbui/protocol"
	"github.com/janpfeifer/gonb/version"
)

//go:generate bash -c "sh version.sh version | tr -d '\n' > version/version.txt"
//go:generate bash -c "sh version.sh hash | tr -d '\n' > version/hash.txt"
func must(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	must(os.Setenv(protocol.GONB_GIT_COMMIT, version.AppVersion.Commit))
	must(os.Setenv(protocol.GONB_VERSION, version.AppVersion.Version))
}
