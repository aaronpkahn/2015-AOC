package main

import (
	"fmt"
)

type deer struct {
	n        string
	s, t, r  int
	flyTime  int
	restTime int
	distance int
	points   int
}

func main() {
	deers := []deer{{n: "Rudolph", s: 22, t: 8, r: 165}, {n: "Cupid", s: 8, t: 17, r: 114}, {n: "Prancer", s: 18, t: 6, r: 103}, {n: "Donner", s: 25, t: 6, r: 145}, {n: "Dasher", s: 11, t: 12, r: 125}, {n: "Comet", s: 21, t: 6, r: 121}, {n: "Blitzen", s: 18, t: 3, r: 50}, {n: "Vixen", s: 20, t: 4, r: 75}, {n: "Dancer", s: 7, t: 20, r: 119}}
	d := fly(2503, deers)
	fmt.Println(d)
	max := 0
	for _, v := range d {
		if v.points > max {
			max = v.points
		}
	}
	fmt.Println(max)
}

func fly(t int, d []deer) []deer {
	max := 0
	for i := 1; i <= t; i++ {
		for j, _ := range d {
			v := &d[j]
			if v.restTime > 0 {
				v.restTime--
				continue
			}
			v.distance += v.s
			if v.distance > max {
				max = v.distance
			}
			v.flyTime++
			if v.flyTime == v.t {
				v.restTime = v.r
				v.flyTime = 0
			}
		}
		for j, _ := range d {
			v := &d[j]
			if v.distance == max {
				v.points++
			}
		}
	}

	fmt.Println(d)
	return d
}
