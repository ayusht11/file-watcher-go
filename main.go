package main

import (
	"github.com/fleek-test-task/file_watcher/cmd"
)

func main() {

	cli := cmd.Cmd{}
	cli.Start_watcher()
}
