package internal

import (
	"fmt"
	"os"
)

var VERSION = "DEV-1.0"
var GOVERSION = "UNKNOWN"
var GITCOMMIT = "UNKNOWN*"
var BUILDTIME = "UNKNOWN"

func version() string {
	return fmt.Sprintf(`CodeMerge %s
Git commit: %s,
Build with %s
Build at %s`, VERSION, GITCOMMIT, GOVERSION, BUILDTIME)
}

func ShowVersion(showVersion *bool) {
	if *showVersion {
		fmt.Println(version())
		os.Exit(0)
	}
}
