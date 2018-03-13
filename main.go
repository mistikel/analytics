package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	Minute    int
	Directory string
)

func main() {
	var Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s[analytics]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.IntVar(&Minute, "t", 0, "Enter minute (required)")
	flag.StringVar(&Directory, "d", "", "Enter directory full path (required)")
	flag.Parse()

	if Minute == 0 && Directory == "" {
		Usage()
		os.Exit(1)
	}

	now := time.Now()

	timeLimit := now.Add(time.Duration(-1*Minute) * time.Minute)

	fmt.Printf("Minute : %d and Directory : %s\n", Minute, Directory)
	fmt.Println("Now :", now)
	fmt.Println("Limit :", timeLimit)

	files, err := ioutil.ReadDir(Directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.ModTime().Before(timeLimit) {
			fmt.Println(file.Name(), file.ModTime())
		}
	}
}
