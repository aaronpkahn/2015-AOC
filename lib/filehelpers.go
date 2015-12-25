package lib

import (
	"bufio"
	"os"
)

func ReadLinesToSlice(path string) (a []string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		a = append(a, scanner.Text())
	}
	return a
}
