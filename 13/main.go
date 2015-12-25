package main

import (
	"adventofcode/lib"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// a := []int{1, 2}
	// fmt.Println(a[:0], a[1:])
	// b := make([]int, len(a))
	// copy(b, a)
	// fmt.Println(append(b[:0], b[1:]...))
	// fmt.Println(a)
	//fmt.Println([]int{1, 2}[:1])

	c := lib.Permutate([]int{0, 1, 2, 3, 4, 5, 6, 7})
	lines := readLinesToSlice("input.txt")
	t := parseLinesToRelationships(lines, 8)
	fmt.Println(len(c))
	max := 0
	for _, v := range c {
		temp := calculateTotal(v, t)
		if temp > max {
			max = temp
		}
	}
	fmt.Println(max)
}

//2+20 15+5 10+1 =

func calculateTotal(o []int, h [][]int) int {
	l := len(o)
	sum := 0 //h[o[0]][o[l-1]] + h[o[l-1]][o[0]]
	for i, v := range o[:l-1] {
		sum += h[v][o[i+1]] + h[o[i+1]][v]
	}
	return sum
}

func parseLinesToRelationships(lines []string, peopleCount int) (r [][]int) {
	//var lastPerson string

	re := regexp.MustCompile(`(\w+) would (\w+) (\d+) happiness units by sitting next to (\w+)`)
	for i := 0; i < len(lines); i++ {
		x := i / (peopleCount - 1)
		y := i % (peopleCount - 1)
		if y == 0 {
			r = append(r, make([]int, 8))
		}

		if y >= x {
			y++
		}

		s := re.FindStringSubmatch(lines[i])
		v, _ := strconv.Atoi(s[3])
		if s[2] == "lose" {
			v = -v
		}
		r[x][y] = v
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
