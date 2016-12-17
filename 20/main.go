// Prime factorization derived from slightly modified version
// of sieve.go in Go source distribution.
package main

import (
	//	"adventofcode/lib"
	"fmt"
)

func main() {
	max := 0
	num := 0
	for i := 811000; i < 850000; i++ {
		t := sumIt(calcAllFactors(i))
		if t > max {
			max = t
			num = i
		}
	}
	fmt.Println(max, num)
}

func sumIt(n []int) (sum int) {
	for _, v := range n {
		sum += v * 10
	}
	return sum
}

func calcAllFactors(n int) (r []int) {
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			r = append(r, i)
		}
	}
	r = append(r, n)
	return r
}
