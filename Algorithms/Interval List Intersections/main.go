package main

import "fmt"

func main() {
	fmt.Println(intervalIntersection([][]int{{4, 11}}, [][]int{{1, 2}, {8, 11}, {12, 13}, {14, 15}, {17, 19}}))
}

func intervalIntersection(A [][]int, B [][]int) [][]int {
	res := make([][]int, 0)
	for i, j := 0, 0; i < len(A) && j < len(B); {
		// A[i] 是 B[j] 的子集
		if A[i][0] >= B[j][0] && A[i][1] <= B[j][1] {
			res = append(res, A[i])
			i++
			continue
		}
		// B[j] 是 A[i] 的子集
		if A[i][0] <= B[j][0] && A[i][1] >= B[j][1] {
			res = append(res, B[j])
			j++
			continue
		}
		// [Bstart, Aend]
		if A[i][0] <= B[j][0] && A[i][1] <= B[j][1] && A[i][1] >= B[j][0] {
			res = append(res, []int{B[j][0], A[i][1]})
			i++
			continue
		}
		// [Astart, Bend]
		if A[i][0] >= B[j][0] && A[i][1] >= B[j][1] && A[i][0] <= B[j][1] {
			res = append(res, []int{A[i][0], B[j][1]})
			j++
			continue
		}
		// 没有交集
		if A[i][1] < B[j][0] {
			i++
			continue
		}
		if A[i][0] > B[j][1] {
			j++
			continue
		}
	}
	return res
}
