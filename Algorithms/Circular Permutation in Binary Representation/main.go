package main

func main() {

}

func circularPermutation(n int, start int) []int {
	ans := make([]int, 1, 1<<n)
	ans[0] = start
	for i := 1; i <= n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, ((ans[j]^start)|(1<<(i-1)))^start)
		}
	}
	return ans
}
