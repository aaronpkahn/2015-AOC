package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	v1()
	//  v2()
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Println(name, "took", elapsed)
}

func v1() {
	defer timeTrack(time.Now(), "v1")
	a := "3113322113"

	for i := 0; i < 50; i++ {
		a = lns(a)
	}
	fmt.Println(len(a))
}

func lns(s string) (r string) {
	j := 1
	a := s[0]
	var c byte
	for i := 1; i < len(s); i++ {
		c = s[i]
		if a == c {
			j += 1
			continue
		}
		r += strconv.Itoa(j) + string(a)
		a = c
		j = 1
	}
	r += strconv.Itoa(j) + string(c)
	return r
}
