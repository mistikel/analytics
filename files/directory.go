package files

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/mistikel/analytics/helper"
)

// GetApproximateFile return int.
// assumpiton : log file generate base on time to each iteration.
func (filesModule *FilesModule) GetApproximateFile(ctx context.Context, path string, limit time.Time) int {
	dir, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	files, err := dir.Readdir(1)
	dir.Close()
	if err != nil {
		log.Fatal(err)
	}

	n := helper.NumberOfFile(files[0].Name())
	var p string
	if n > 9 {
		p = "http-" + strconv.Itoa(n) + ".log"
	} else {
		p = "http-0" + strconv.Itoa(n) + ".log"
	}

	file, _ := os.Stat(path + p)
	x := file.ModTime().Minute()
	b := x - files[0].ModTime().Minute()
	a := helper.GetA(x, b, n)

	appxFile := helper.GetN(limit.Minute(), a, b)

	return appxFile
}

// ReadDirectory return File Info
// get list file in directory
func (filesModule *FilesModule) ReadDirectory(ctx context.Context, path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	return files
}
