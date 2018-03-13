package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	Time      int
	Directory string
)

func main() {
	var Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s[analytics]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.IntVar(&Time, "t", 0, "Enter minute (required)")
	flag.StringVar(&Directory, "d", "", "Enter directory full path (required)")
	flag.Parse()

	if Time == 0 && Directory == "" {
		Usage()
		os.Exit(1)
	}

	fmt.Printf("Time : %d and Directory : %s", Time, Directory)
}
