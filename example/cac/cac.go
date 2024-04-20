package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/txthinkinginc/cac"
)

func init() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "-h", "--help", "help":
			defer func() {
				flag.Usage()
				os.Exit(0)
			}()
			goto stdflag
		}

		fileContent, err := os.ReadFile(os.Args[1])
		if err != nil {
			log.Fatalln(err)
			return
		}

		flag.CommandLine.Parse(cac.Parse(string(fileContent)))
		return
	}

stdflag:
	flag.Parse()
}

func main() {
	fmt.Println(flag.Args())
}
