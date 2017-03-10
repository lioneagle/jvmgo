package cmds

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	Help      bool
	Version   bool
	ClassPath string
	Xjre      string
	Class     string
	Args      []string
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

func ParseCmd() *Cmd {
	cmd := &Cmd{}
	//flag.Usage = PrintUsage
	flag.BoolVar(&cmd.Help, "help", false, "print help message")
	flag.BoolVar(&cmd.Help, "?", false, "print help message")
	flag.BoolVar(&cmd.Version, "version", false, "print version and exit")
	flag.BoolVar(&cmd.Version, "v", false, "print version and exit")
	flag.StringVar(&cmd.Xjre, "Xjre", "", "path to jre")

	flag.StringVar(&cmd.ClassPath, "classpath", "", "classpath")
	flag.StringVar(&cmd.ClassPath, "cp", "", "classpath")
	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		cmd.Class = args[0]
		cmd.Args = args[1:]
	}

	return cmd
}
