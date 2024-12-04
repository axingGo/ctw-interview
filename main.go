package main

import (
	"ctx-interview/cmd"
	"ctx-interview/conf"
)

func init() {
	conf.LoadConfig()
}

func main() {
	err := cmd.CmdExec()
	if err != nil {
		panic(err)
	}
}
