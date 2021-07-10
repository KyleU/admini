package main

import (
	"os"

	"github.com/kyleu/admini/app"
	"github.com/kyleu/admini/app/cmd"
	"github.com/kyleu/admini/app/log"
)

var (
	version = "0.1.4" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.AppBuildInfo = &app.BuildInfo{Version: version, Commit: commit, Date: date}
	logger, err := cmd.Run()
	if err != nil {
		msg := "exiting due to error"
		if logger == nil {
			println(log.Red.Add(err.Error())) //nolint
			println(log.Red.Add(msg))         //nolint
		} else {
			logger.Error(err)
			logger.Error(msg)
		}
		os.Exit(1)
	}
}
