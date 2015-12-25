package main

import (
	"adventofcode/lib"
	"fmt"
	"regexp"
	"strconv"
)

type instruction struct {
	X0, Y0, X1, Y1 int
	Action         string
}

func main() {
	lines := lib.ReadLinesToSlice("input.txt")
	inst := parseInstructions(lines)
	var grid [1000][1000]int
	for _, v := range inst {
		executeCommand(&grid, v)
	}
	fmt.Println(countOn(grid))
}

func modit(b *[2][2]int) {
	b[1][1] = 5
}

func countOn(grid [1000][1000]int) (sum int) {
	for _, v := range grid {
		for _, iv := range v {
			sum += iv
		}
	}
	return sum
}

func executeCommand(grid *[1000][1000]int, ins instruction) {
	for i := ins.X0; i <= ins.X1; i++ {
		for j := ins.Y0; j <= ins.Y1; j++ {
			switch ins.Action {
			case "turn on":
				grid[i][j] += 1
			case "turn off":
				if grid[i][j] > 0 {
					grid[i][j] -= 1
				}
			case "toggle":
				grid[i][j] += 2
			}
		}
	}
}

func parseInstructions(lines []string) (ia []instruction) {
	re := regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)
	for _, v := range lines {
		s := re.FindAllStringSubmatch(v, -1)
		action := s[0][1]
		x0, _ := strconv.Atoi(s[0][2])
		y0, _ := strconv.Atoi(s[0][3])
		x1, _ := strconv.Atoi(s[0][4])
		y1, _ := strconv.Atoi(s[0][5])
		ia = append(ia, instruction{X0: x0, Y0: y0, X1: x1, Y1: y1, Action: action})
	}
	return ia
}
