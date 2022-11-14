package main

import "fmt"

func main() {
	fmt.Println(orderOfLargestPlusSign(5, [][]int{{4, 2}}))
	fmt.Println(orderOfLargestPlusSign(1, [][]int{{0, 0}}))
	fmt.Println(orderOfLargestPlusSign(2, [][]int{{0, 0}, {0, 1}, {1, 0}}))
	fmt.Println(orderOfLargestPlusSign(2, [][]int{{0, 0}, {1, 1}}))
	fmt.Println(orderOfLargestPlusSign(10, [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 7}, {1, 2}, {1, 3}, {1, 9}, {2, 3}, {2, 5}, {2, 7}, {2, 8}, {3, 2}, {3, 5}, {3, 7}, {4, 2}, {4, 3}, {4, 5}, {4, 7}, {5, 1}, {5, 4}, {5, 8}, {5, 9}, {7, 2}, {7, 5}, {7, 7}, {7, 8}, {8, 5}, {8, 8}, {9, 0}, {9, 1}, {9, 2}, {9, 8}}))
}

type Point struct {
	Up   int
	Left int
}

func orderOfLargestPlusSign(n int, mines [][]int) int {
	mMines := make(map[string]struct{})
	for _, v := range mines {
		mMines[key(v[0], v[1])] = struct{}{}
	}
	isZero := func(x, y int) bool {
		_, ok := mMines[key(x, y)]
		return ok
	}
	dp := make([][]Point, n)
	for x := 0; x < n; x++ {
		dp[x] = make([]Point, n)
		for y := 0; y < n; y++ {
			if !isZero(x, y) {
				dp[x][y].Up = 1
				if x-1 >= 0 {
					dp[x][y].Up = dp[x-1][y].Up + 1
				}
				dp[x][y].Left = 1
				if y-1 >= 0 {
					dp[x][y].Left = dp[x][y-1].Left + 1
				}
			}
		}
	}
	// for _, i := range dp {
	// 	for _, v := range i {
	// 		fmt.Printf("(Up:%d,Left:%d) ", v.Up, v.Left)
	// 	}
	// 	fmt.Println()
	// }
	ans := 0
	for x := 0; x < len(dp); x++ {
		for y := 0; y < len(dp[x]); y++ {
			minK := min(dp[x][y].Left, dp[x][y].Up)
			for i := minK - 1; i >= 0; i-- {
				if x+i > len(dp)-1 || y+i > len(dp[x])-1 {
					continue
				}
				if dp[x+i][y].Up-dp[x][y].Up >= i && dp[x][y+i].Left-dp[x][y].Left >= i {
					if i+1 > ans {
						ans = i + 1
					}
					break
				}
			}
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func key(y, x int) string {
	return fmt.Sprintf("%d-%d", y, x)
}
