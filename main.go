package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

func main() {
	if runtime.GOOS != "windows" {
		fmt.Println("Must be run on windows")
		os.Exit(1)
	}

	symbolicFlag := flag.Bool("s", false, "Create a symbolic link")
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("LINK and TARGET are required")
		os.Exit(1)
	}

	target := args[0]
	link := args[1]

	var err error
	if *symbolicFlag {
		err = os.Symlink(target, link)
	} else {
		os.Link(target, link)
	}

	if err != nil {
		panic(err)
	}
}
