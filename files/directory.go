package files

import (
	"context"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/mistikel/analytics/helper"
)

// GetApproximateFile return int.
// assumpiton : log file generate base on time to each iteration.
func (filesModule *FilesModule) GetApproximateFile(ctx context.Context, path string, limit time.Time) (int, error) {
	dir, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	files, err := dir.Readdir(1)
	dir.Close()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	n := helper.NumberOfFile(files[0].Name())
	var p string
	if n > 9 {
		p = "http-" + strconv.Itoa(n) + ".log"
	} else {
		p = "http-0" + strconv.Itoa(n) + ".log"
	}

	file, _ := os.Stat(path + p)
	x := file.ModTime().UTC().Minute()
	b := x - files[0].ModTime().UTC().Minute()
	a := helper.GetA(x, b, n)

	if b == 0 {
		log.Fatal("Can't work with 0")
		return 0, errors.New("can't work with 0")
	}
	appxFile := helper.GetN(limit.UTC().Minute(), a, b)

	return appxFile, nil
}

// ReadDirectory return File Info
// get list file in directory
func (filesModule *FilesModule) ReadDirectory(ctx context.Context, path string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return files, err
	}

	return files, nil
}
