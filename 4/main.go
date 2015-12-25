package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	s := "iwrupvqb"
	i := 0
	for {
		x := md5.Sum([]byte(s + strconv.Itoa(i)))
		t := fmt.Sprintf("%x", x)
		if t[:6] == "000000" {
			break
		}
		i++
	}
	fmt.Println(i)

	fmt.Printf("%x", md5.Sum([]byte("iwrupvqb9958218")))
}
