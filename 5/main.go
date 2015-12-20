package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {

	lines := readLinesToSlice("input.txt")
	count := 0
	for _, l := range lines {
		if norn2(l) {
			count++
		}
	}
	fmt.Println(count)
}

func norn2(s string) bool {
	twoLetters := false
	letterPaired := false
	m := make(map[string]int)
	for i, l := range s {
		if !twoLetters && i > 1 {
			if l == rune(s[i-2]) {
				twoLetters = true
			}
		}
		if !letterPaired && i > 0 {
			pair := string(s[i-1]) + string(l)
			if v, ok := m[pair]; ok {
				if v+2 < i {
					letterPaired = true
				}
			} else {
				m[pair] = i - 1
			}
		}
	}
	return twoLetters && letterPaired
}

var r = regexp.MustCompile("ab|cd|pq|xy")

func norn1(s string) bool {
	if r.MatchString(s) {
		return false
	}
	vowelCount := 0
	twoLetters := false
	lastLetter := '@'

	for _, i := range s {
		if !twoLetters {
			if lastLetter == i {
				twoLetters = true
			}
			lastLetter = i
		}
		if i == 'a' || i == 'e' || i == 'i' || i == 'o' || i == 'u' {
			vowelCount++
		}
	}
	return vowelCount >= 3 && twoLetters
}

func readLinesToSlice(path string) (a []string) {
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		a = append(a, scanner.Text())
	}
	return a
}
