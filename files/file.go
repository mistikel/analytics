package files

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"
)

// ReadFile return nil or error
// Read single file line by line
func (filesModule *FilesModule) ReadFile(ctx context.Context, path string, limit time.Time, wg *sync.WaitGroup) error {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Can not find file: ", path)
		return err
	}

	defer wg.Done()

	reader := bufio.NewReader(file)

	text, _, err := reader.ReadLine()
	for err == nil {
		reg := regexp.MustCompile("\\[.*\\]")
		res := reg.FindAllString(string(text), -1)
		if len(res) > 0 {
			t := res[0][1:3] + " " + res[0][4:7] + " " + res[0][10:12] + " " + res[0][13:15] + ":" + res[0][16:18] + " MST"
			date, err := time.Parse(time.RFC822, t)
			if err != nil {
				break
			}

			//check if date is before time limit
			if date.Before(limit) {
				break
			}
			fmt.Println(string(text))
		}
		text, _, err = reader.ReadLine()
	}

	return nil
}
