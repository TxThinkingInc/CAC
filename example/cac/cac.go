package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/txthinkinginc/cac"
)

var (
	commandLine = flag.NewFlagSet("cac", flag.ExitOnError)

	showVersion bool
	foo         string
)

func init() {
	commandLine.BoolVar(&showVersion, "v", false, "Show version")
	commandLine.StringVar(&foo, "foo", "", "Some argument")
}

func main() {
	var args []string
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "-h", "--help", "help":
			commandLine.Usage()
			return
		default:
			fileContent, err := os.ReadFile(os.Args[1])
			if err != nil {
				args = os.Args[1:]
			} else {
				args = cac.Parse(string(fileContent))
			}
		}
	}

	err := commandLine.Parse(args)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(commandLine.Args())
	fmt.Println("showVersion: ", showVersion)
	fmt.Println("foo: ", foo)
}
