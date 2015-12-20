package main

import (
	"fmt"
)

func main() {
	a := []rune("cqjxxzaa")
	fmt.Println(check(a))
	for {
		if check(a) {
			break
		}
		a = increment(a)
	}
	fmt.Println(string(a))
}

func check(r []rune) bool {
	straight := false
	pairs := false
	p := make([]rune, 0)

	for i, v := range r {
		if v == 'i' || v == 'o' || v == 'l' {
			return false
		}
		if !straight && i > 1 {
			if r[i-2] == v-2 && r[i-1] == v-1 {
				straight = true
			}
		}

		if !pairs && i > 0 {
			if v == r[i-1] {
				p = append(p, v)
				if len(p) > 2 || p[0] != v {
					pairs = true
				}
			}
		}

	}
	return pairs && straight
}

func increment(r []rune) []rune {
	for i := len(r) - 1; i >= 0; i-- {
		c := r[i]
		r[i] = (c-97+1)%26 + 97
		if c != 'z' {
			break
		}
	}
	return r
}
