package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	lines := readLinesToSlice("input.txt")
	parseLinesToRelationships(lines, 8)
}

func parseLinesToRelationships(lines []string, peopleCount int) (r [][]int) {
	//var lastPerson string

	re := regexp.MustCompile(`(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+)`)
	for i, j := 0, 0; i < len(lines); i++ {
		if i%peopleCount == 0 {
			j++
		}
		i % peopleCount
		fmt.Println(re.FindStringSubmatch(lines[i]))

	}
	return r
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
