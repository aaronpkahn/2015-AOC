package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`(\d+)x(\d+)x(\d+)`)

func main() {
	lines := readLinesToSlice("input.txt")
	//fmt.Println(re.FindStringSubmatch(lines[0])[1:4])
	p, r := findSum(lines)
	fmt.Println(p+r, p, r)
}
func findSum(lines []string) (paper int, ribbon int) {
	for _, line := range lines {
		lwh := re.FindStringSubmatch(line)[1:4]
		l, _ := strconv.Atoi(lwh[0])
		w, _ := strconv.Atoi(lwh[1])
		h, _ := strconv.Atoi(lwh[2])
		lw := l * w
		lh := l * h
		wh := w * h
		paper += 2*lw + 2*lh + 2*wh + theMin(lw, lh, wh)
		lwp := 2*l + 2*w
		lhp := 2*l + 2*h
		whp := 2*w + 2*h
		ribbon += theMin(lwp, lhp, whp) + l*w*h
	}
	return paper, ribbon
}

func theMin(a ...int) int {
	min := a[0]
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
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
