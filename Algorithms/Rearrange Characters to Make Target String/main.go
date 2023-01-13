package main

import "math"

func main() {

}

func rearrangeCharacters(s string, target string) int {
	ms := make(map[int32]int)
	mt := make(map[int32]int)
	for _, b := range s {
		ms[b]++
	}
	for _, b := range target {
		mt[b]++
	}
	ans := math.MaxInt
	for b, c := range mt {
		v, ok := ms[b]
		if !ok {
			return 0
		}
		x := v / c
		if x < ans {
			ans = x
		}
	}
	return ans
}
