package main

import (
	"cmds"
	"fmt"
)

func main() {

	cmd := cmds.ParseCmd()
	if cmd.Version {
		fmt.Println("version 0.0.1")
	} else if cmd.Help {
		cmds.PrintUsage()
	} else {
		startJVM(cmd)
	}

	//fmt.Println("cmd =", cmd)

}

func startJVM(cmd *cmds.Cmd) {
	fmt.Printf("classpath: %s, class: %s, args: %v\n", cmd.ClassPath, cmd.Class, cmd.Args)
}
