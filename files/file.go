package files

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
)

type chunk struct {
	bufsize int
	offset  int64
}

// ReadFile return nil or error
// Read single file line by line
func (filesModule *FilesModule) ReadFile(ctx context.Context, path string, wg *sync.WaitGroup) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Can not find file: ", path)
		return
	}

	defer wg.Done()

	reader := bufio.NewReader(file)

	text, _, err := reader.ReadLine()
	for err == nil {

		fmt.Println(file.Name(), string(text))
		text, _, err = reader.ReadLine()
	}
}
