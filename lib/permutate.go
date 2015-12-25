package lib

func Permutate(a []int) (r [][]int) {
	return mix(a)
}

func mix(a []int) (r [][]int) {
	if len(a) == 1 {
		return append(r, a)
	}
	for i, v := range a {
		b := make([]int, len(a))
		copy(b, a)
		ns := []int{v}
		for _, x := range mix(append(b[:i], b[i+1:]...)) {
			r = append(r, append(ns, x...))
		}
	}
	return r
}
