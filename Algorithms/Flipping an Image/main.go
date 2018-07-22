package main

import "fmt"

func main() {

	A := [][]int{{1,1,0,0},{1,0,0,1},{0,1,1,1},{1,0,1,0}}
	fmt.Println(flipAndInvertImage(A))
}

func flipAndInvertImage(A [][]int) [][]int {

	for Ai, a := range A {
		for ai, _ := range a {
			if len(a) % 2 == 0 {
				if ai >= len(a)/2 {
					break
				}
			}else{
				if ai > len(a)/2 {
					break
				}
			}
			A[Ai][ai], A[Ai][len(a)-1-ai] = 1-A[Ai][len(a)-1-ai], 1-A[Ai][ai]
		}
	}
	return A
}
