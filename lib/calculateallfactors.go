// Prime factorization derived from slightly modified version
// of sieve.go in Go source distribution.
package lib

// Generate numbers until the limit max.
// after the 2, all the prime numbers are odd
// Send a channel signal when the limit is reached
func generate(max int, ch chan<- int) {
	ch <- 2
	for i := 3; i <= max; i += 1 {
		ch <- i
	}
	ch <- -1 // signal that the limit is reached
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan int, out chan<- int, prime int) {
	for i := <-in; i != -1; i = <-in {
		//if i%prime != 0 {
		out <- i
		//}
	}
	out <- -1
}

func CalcFactors(number_to_factorize int) []int {
	rv := []int{}
	ch := make(chan int)
	go generate(number_to_factorize, ch)
	for prime := <-ch; (prime != -1) && (number_to_factorize > 1); prime = <-ch {
		for number_to_factorize%prime == 0 {
			number_to_factorize = number_to_factorize / prime
			rv = append(rv, prime)
		}
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	return rv
}
