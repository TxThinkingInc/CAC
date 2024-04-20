package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/txthinkinginc/cac"
)

var flagSet *flag.FlagSet

func init() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "-h", "--help", "help":
			goto stdParse
		}

		fileContent, err := os.ReadFile(os.Args[1])
		if err != nil {
			log.Fatalln(err)
			return
		}

		flagSet = flag.NewFlagSet("custom flag set", flag.ExitOnError)
		flagSet.Usage = func() {
			fmt.Println("Help of custom flag set...")
		}
		_ = flagSet.Parse(cac.Parse(string(fileContent)))
		return
	}

stdParse:
	flag.Parse()
}

func main() {
	fmt.Println(flagSet.Args())
}
