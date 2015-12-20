package main

import (
	"fmt"
	"regexp"
	"sort"
)

var mapping = [][]string{{"Al", "ThF"}, {"Al", "ThRnFAr"}, {"B", "BCa"}, {"B", "TiB"}, {"B", "TiRnFAr"}, {"Ca", "CaCa"}, {"Ca", "PB"}, {"Ca", "PRnFAr"}, {"Ca", "SiRnFYFAr"}, {"Ca", "SiRnMgAr"}, {"Ca", "SiTh"}, {"F", "CaF"}, {"F", "PMg"}, {"F", "SiAl"}, {"H", "CRnAlAr"}, {"H", "CRnFYFYFAr"}, {"H", "CRnFYMgAr"}, {"H", "CRnMgYFAr"}, {"H", "HCa"}, {"H", "NRnFYFAr"}, {"H", "NRnMgAr"}, {"H", "NTh"}, {"H", "OB"}, {"H", "ORnFAr"}, {"Mg", "BF"}, {"Mg", "TiMg"}, {"N", "CRnFAr"}, {"N", "HSi"}, {"O", "CRnFYFAr"}, {"O", "CRnMgAr"}, {"O", "HP"}, {"O", "NRnFAr"}, {"O", "OTi"}, {"P", "CaP"}, {"P", "PTi"}, {"P", "SiRnFAr"}, {"Si", "CaSi"}, {"Th", "ThCa"}, {"Ti", "BP"}, {"Ti", "TiTi"}, {"e", "HF"}, {"e", "NAl"}, {"e", "OMg"}}
var start = "CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr"

//var mapping = [][]string{{"e", "H"}, {"e", "O"}, {"H", "HO"}, {"H", "OH"}, {"O", "HH"}}
//var start = "HOHOHO"

type BySize [][]string

func (a BySize) Len() int           { return len(a) }
func (a BySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySize) Less(i, j int) bool { return len(a[i][1])-len(a[i][0]) > len(a[j][1])-len(a[j][0]) }

func main() {
	// start = Reverse(start)
	// for _, v := range mapping {
	// 	v[0] = Reverse(v[0])
	// 	v[1] = Reverse(v[1])
	// }
	sort.Sort(BySize(mapping))
	fmt.Println(mapping)
	fmt.Println(start)
	fmt.Println(incrementDown(start))
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func incrementDown(s string) int {
	steps := 0
	var re = make([]*regexp.Regexp, len(mapping))
	for i, v := range mapping {
		re[i] = regexp.MustCompile(v[1])
	}
	for {
		for i, v := range mapping {
			id := re[i].FindStringIndex(s)
			if id == nil {
				if i == len(mapping)-1 {
					return -1
				}
				continue
			}
			//fmt.Println(id)
			s = s[:id[0]] + v[0] + s[id[1]:]
			fmt.Println(s)
			//fmt.Println("--------")
			steps++
			if s == "e" {
				return steps
			}
			break
		}
	}
	return -1
}

func countDistinctChanges(s string) int {
	m := make(map[string]bool)
	for _, sp := range mapping {
		re := regexp.MustCompile(sp[0])
		for _, id := range re.FindAllStringSubmatchIndex(s, -1) {
			m[s[:id[0]]+sp[1]+s[id[1]:]] = true
		}
	}
	return len(m)
}
