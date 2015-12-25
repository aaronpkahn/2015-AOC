package main

import (
	"adventofcode/lib"
	"fmt"
	"strconv"
)

func main() {
	//a := "\xa8br\x8bjr\""
	lines := lib.ReadLinesToSlice("input.txt")
	code, val := 0, 0
	for _, v := range lines {
		code += len(v)
		s, _ := strconv.Unquote(v)
		val += len(s)
	}
	fmt.Println(code, val, code-val)
	lit, quo := 0, 0
	for _, v := range lines {
		lit += len(v)
		s := strconv.Quote(v)
		quo += len(s)
	}
	fmt.Println(lit, quo, quo-lit)

	//fmt.Printf("%s", lines[0])
}
