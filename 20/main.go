// Prime factorization derived from slightly modified version
// of sieve.go in Go source distribution.
package main

import (
	"adventofcode/20/lib"
	"fmt"
)

func main() {
	fmt.Println(lib.CalcPrimeFactors(1000000))
}
