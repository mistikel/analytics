package files

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestReadFile(t *testing.T) {
	fm := NewFilesModule()
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(1)
	res := fm.ReadFile(ctx, "../logs/", time.Now(), &wg)
	if res != nil {
		t.Errorf("Expected nil got %s", res)
	}
	wg.Wait()
}
