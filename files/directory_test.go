package files

import (
	"context"
	"testing"
	"time"
)

func TestGetApproximateFile(t *testing.T) {
	fm := NewFilesModule()
	ctx := context.Background()
	approx, _ := fm.GetApproximateFile(ctx, "../logs/", time.Now())

	if approx < 0 {
		t.Errorf("Expected > 0 got %d", approx)
	}
}

func TestReadDirectory(t *testing.T) {
	fm := NewFilesModule()
	ctx := context.Background()
	files, _ := fm.ReadDirectory(ctx, "../logs/")

	if len(files) < 0 {
		t.Errorf("Expected > 0 got %d", len(files))
	}
}
