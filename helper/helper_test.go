package helper

import (
	"testing"
)

func TestNumberOfFile(t *testing.T) {
	file := "http-02.log"
	n := NumberOfFile(file)
	if n != 3 {
		t.Errorf("Expected 3 got %d", n)
	}

}

func TestGetA(t *testing.T) {
	a := GetA(6, 2, 3)
	if a != 2 {
		t.Errorf("Expected 2 got %d", a)
	}
}

func TestGetN(t *testing.T) {
	n := GetN(6, 2, 2)
	if n != 3 {
		t.Errorf("Expected 3 got %d", n)
	}
}
