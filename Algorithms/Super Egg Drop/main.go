package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(superEggDrop(1, 2))
	fmt.Println(superEggDrop(2, 6))

	s := time.Now()
	fmt.Println(superEggDrop(4, 5000))
	fmt.Println(time.Since(s))

	s1 := time.Now()
	fmt.Println(superEggDrop2(4, 5000))
	fmt.Println(time.Since(s1))

	s2 := time.Now()
	fmt.Println(superEggDrop3(4, 5000))
	fmt.Println(time.Since(s2))
}

func superEggDrop(K int, N int) int {
	memo := make(map[string]int)
	return dp(K, N, memo)
}

func superEggDrop2(K int, N int) int {
	memo := make(map[string]int)
	return dp2(K, N, memo)
}

func superEggDrop3(K int, N int) int {
	dp := make([][]int, K+1)
	for i := 0; i <= K; i++ {
		dp[i] = make([]int, N+1)
	}

	m := 0
	for dp[K][m] < N {
		m++

		for i := 1; i <= K; i++ {
			dp[i][m] = dp[i][m-1] + dp[i-1][m-1] + 1
		}
	}
	return m
}

func dp(K, N int, memo map[string]int) int {
	if K == 1 {
		return N
	}
	if N == 0 {
		return 0
	}
	if v, ok := memo[fmt.Sprintf("%d%d", K, N)]; ok {
		return v
	}

	res := N + 1
	for i := 1; i <= N; i++ {
		res = min(res, max(dp(K, N-i, memo), dp(K-1, i-1, memo))+1)
	}
	memo[fmt.Sprintf("%d%d", K, N)] = res

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dp2(K, N int, memo map[string]int) int {
	if K == 1 {
		return N
	}
	if N == 0 {
		return 0
	}
	if v, ok := memo[fmt.Sprintf("%d%d", K, N)]; ok {
		return v
	}

	res := N + 1
	left, right := 0, N
	for left <= right {
		mid := left + (right-left)/2
		broken := dp(K-1, mid-1, memo)
		notBroken := dp(K, N-mid, memo)
		if broken > notBroken {
			right = mid - 1
			res = min(res, broken+1)
		} else {
			left = mid + 1
			res = min(res, notBroken+1)
		}
	}
	memo[fmt.Sprintf("%d%d", K, N)] = res

	return res
}
