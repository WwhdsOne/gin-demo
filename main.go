package main

import (
	"gin-demo/cmd"
	"gin-demo/core"
)

func main() {
	cmd.RootCmd.Execute()
	core.RunServer()
}
