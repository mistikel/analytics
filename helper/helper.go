package helper

import (
	"log"
	"regexp"
	"strconv"
)

// NumberOfFile return int.
// get number in filename.
func NumberOfFile(file string) int {
	reg := regexp.MustCompile("[0-9]+")

	f := reg.FindAllString(file, -1)
	if len(f) < 1 {
		log.Fatal("There is no http-xx.log file")
	}

	n, _ := strconv.Atoi(f[0])

	return n + 1
}

// GetA return int.
// get first value in aritmathic,
func GetA(x, b, n int) int {
	return x - b*(n-1)
}

// GetN return int,
// get n-value in aritmathic.
func GetN(x, a, b int) int {
	return ((x - a) / b) + 1
}
