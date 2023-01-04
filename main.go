/*
Copyright © 2023 jaronnie jaron@jaronnie.com

*/

package main

import (
	"github.com/rdsutbbp/logx"

	"github.com/jaronnie/gpm-runtime/v2/cmd"
)

var (
	version string
	commit  string
)

func main() {
	logx.Logger()

	cmd.Version = version
	cmd.Commit = commit

	cmd.Execute()
}
