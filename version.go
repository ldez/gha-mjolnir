package main

import (
	"fmt"
	"runtime"
)

var (
	version = "dev"
	commit  = "I don't remember exactly"
	date    = "I don't remember exactly"
)

// displayVersion DisplayVersion version.
func displayVersion() {
	fmt.Printf(`mjolnir:
 version     : %s
 commit      : %s
 build date  : %s
 go version  : %s
 go compiler : %s
 platform    : %s/%s
`, version, commit, date, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)
}
