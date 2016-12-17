package main

import (
	"fmt"
)

func main() {
	n := []int{1, 3, 5, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113}
	sum := 0
	for _, v := range n {
		sum += v
	}

	fmt.Println(findBalance(sum/3, n))
}

func findBalance(goal int, n []int) (r [][]int) {
	for i := 0; i < 3; i++ {

		append()
	}
	// for i := len(n)-1; i>=0; i--{

	// }
}
func findSolution(goal int, n []int) (r []int) {

}
