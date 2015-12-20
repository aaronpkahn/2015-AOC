package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile("(-{0,1}[1-9][0-9]*)")

func main() {
	s, _ := ioutil.ReadFile("output.txt")
	a := string(s)
	fmt.Println(len(a))
	sum := 0
	for _, v := range re.FindAllStringSubmatch(a, -1) {
		iv, _ := strconv.Atoi(v[0])
		sum += iv
		fmt.Println(iv, sum)
	}
	fmt.Println(sum)
}
