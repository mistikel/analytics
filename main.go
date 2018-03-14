package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/mistikel/analytics/files"
)

var (
	Minute    int
	Directory string
	Huge      bool
)

func main() {
	var Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s[analytics]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.IntVar(&Minute, "t", 0, "Enter minute (required)")
	flag.StringVar(&Directory, "d", "", "Enter directory full path (required)")
	flag.BoolVar(&Huge, "b", false, "Use this if you want process big directory (optional)")
	flag.Parse()

	if Minute == 0 && Directory == "" {
		Usage()
		os.Exit(1)
	}

	ctx := context.Background()
	run(ctx)
}

func run(ctx context.Context) {
	now := time.Now()
	timeLimit := now.Add(time.Duration(-1*Minute) * time.Minute)

	fileModule := files.NewFilesModule()

	if Huge {
		number := fileModule.GetApproximateFile(ctx, Directory, timeLimit)
		fmt.Println(number)
	} else {
		files := fileModule.ReadDirectory(ctx, Directory)
		path := Directory + files[0].Name()
		fileModule.ReadFile(ctx, path)
	}
}
