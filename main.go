package main

import "ctx-interview/cmd"

func main() {
	err := cmd.CmdExec()
	if err != nil {
		panic(err)
	}
}
