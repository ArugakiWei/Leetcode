package main

func main() {

}

func dieSimulator(n int, rollMax []int) int {
	const mod int = 1e9 + 7
	const m = 6
	f := make([][m]int, n)
	for j := range f[0] {
		f[0][j] = 1
	}
	s := make([]int, n)
	s[0] = m
	for i := 1; i < n; i++ {
		for j, mx := range rollMax {
			res := s[i-1]
			if i > mx {
				res -= s[i-mx-1] - f[i-mx-1][j]
			} else if i == mx {
				res--
			}
			f[i][j] = (res%mod + mod) % mod // 防止出现负数
			s[i] = (s[i] + f[i][j]) % mod
		}
	}
	return s[n-1]
}
