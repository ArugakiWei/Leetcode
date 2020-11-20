package main

import "fmt"

func main() {
	fmt.Println(stoneGame([]int{5, 3, 4, 5}))
}

type Score struct {
	First  int
	Second int
}

func stoneGame(piles []int) bool {
	dp := make([][]Score, len(piles))
	for i, _ := range dp {
		dp[i] = make([]Score, len(piles))
		dp[i][i] = Score{
			First:  piles[i],
			Second: 0,
		}
	}

	for l := 2; l <= len(piles); l++ {
		for i := 0; i <= len(piles)-l; i++ {
			j := l + i - 1

			/*
				dp[i][j].fir = max(    选择最左边的石头堆     ,     选择最右边的石头堆     )
				解释：我作为先手，面对 piles[i...j] 时，有两种选择：
				要么我选择最左边的那一堆石头，然后面对 piles[i+1...j]
				但是此时轮到对方，相当于我变成了后手；
				要么我选择最右边的那一堆石头，然后面对 piles[i...j-1]
				但是此时轮到对方，相当于我变成了后手。
			*/
			left := piles[i] + dp[i+1][j].Second
			right := piles[j] + dp[i][j-1].Second

			dp[i][j].First = max(left, right)

			/*
				解释：我作为后手，要等先手先选择，有两种情况：
				如果先手选择了最左边那堆，给我剩下了 piles[i+1...j]
				此时轮到我，我变成了先手；
				如果先手选择了最右边那堆，给我剩下了 piles[i...j-1]
				此时轮到我，我变成了先手。
			*/
			if left > right {
				dp[i][j].Second = dp[i+1][j].First
			} else {
				dp[i][j].Second = dp[i][j-1].First
			}
		}
	}

	return dp[0][len(piles)-1].First > dp[0][len(piles)-1].Second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
