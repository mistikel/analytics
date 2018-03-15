package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/mistikel/analytics/files"
)

var (
	Minute    string
	Directory string
	Huge      bool
	wg        sync.WaitGroup
)

func main() {
	var Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [analytics -t <mins>m -d <dir>]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.StringVar(&Minute, "t", "0", "Enter minute followed with 'm' (required)")
	flag.StringVar(&Directory, "d", "", "Enter directory full path ended with '/' (required)")
	flag.BoolVar(&Huge, "b", false, "Use this if you want process big directory (optional)")
	flag.Parse()

	if Minute == "0" && Directory == "" {
		Usage()
		os.Exit(1)
	}

	ctx := context.Background()
	run(ctx)
}

func run(ctx context.Context) {
	now := time.Now().UTC()
	val, err := strconv.Atoi(Minute)
	if end := len(Minute) - 1; end >= 0 && Minute[end] == 'm' && err != nil {
		val, err = strconv.Atoi(Minute[:end])

		if err != nil {
			fmt.Println("Minutes not recognize")
			return
		}
	}

	timeLimit := now.Add(time.Duration(-1*val) * time.Minute)

	fileModule := files.NewFilesModule()

	var path string

	if Huge {
		number, err := fileModule.GetApproximateFile(ctx, Directory, timeLimit)
		if err != nil {
			return
		}

		for number > 0 {
			if number > 9 {
				path = "http-" + strconv.Itoa(number) + ".log"
			} else {
				path = "http-0" + strconv.Itoa(number) + ".log"
			}

			wg.Add(1)

			// go concurrent
			go fileModule.ReadFile(ctx, path, timeLimit, &wg)
		}
	} else {
		files, err := fileModule.ReadDirectory(ctx, Directory)
		if err != nil {
			return
		}

		for _, file := range files {
			if file.ModTime().Before(timeLimit) {
				break
			}
			wg.Add(1)
			path = Directory + file.Name()

			// go concurrent
			go fileModule.ReadFile(ctx, path, timeLimit, &wg)

		}
	}

	wg.Wait()
}
