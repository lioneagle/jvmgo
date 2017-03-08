package main

import (
	"classpath"
	"cmds"
	"fmt"
	"strings"
)

func main() {

	cmd := cmds.ParseCmd()
	if cmd.Version {
		fmt.Println("version 0.0.1")
	} else if cmd.Help || cmd.Class == "" {
		cmds.PrintUsage()
	} else {
		startJVM(cmd)
	}

	fmt.Println("cmd =", cmd)

}

func startJVM(cmd *cmds.Cmd) {
	classPath := classpath.Parse(cmd.Xjre, cmd.ClassPath)
	fmt.Printf("classpath: %s, class: %v, args: %v\n", classPath, cmd.Class, cmd.Args)
	className := strings.Replace(cmd.Class, ".", "/", -1)
	classData, _, err := classPath.ReadClass(className)
	if err != nil {
		fmt.Printf("Can not find or load main class %s\n", cmd.Class)
		return
	}
	fmt.Printf("class data: %v\n", classData)
}
