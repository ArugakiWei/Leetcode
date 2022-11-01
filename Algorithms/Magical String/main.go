package main

import "fmt"

func main() {
	fmt.Println(magicalString(6))
	fmt.Println(magicalString(5))
}

func magicalString(n int) int {
	if n <= 3 {
		return 1
	}
	ms, cm := make([]byte, n+1), make([]byte, n)
	ms[0], cm[0] = '1', '1'
	ms[1], cm[1] = '2', '2'
	ms[2] = '2'
	msSum, cmSum := 3, 3
	for i, j := 3, 2; i < n; {
		d := cmSum - msSum
		if d == 2 {
			ms[i], ms[i+1] = diff(ms[i-1]), diff(ms[i-1])
			msSum += 2
			i += 2
		} else if d == 1 {
			ms[i] = diff(ms[i-1])
			msSum += 1
			i++
		} else {
			cm[j] = ms[j]
			cmSum += count(cm[j])
			j++
		}
	}
	ans := 0
	for i := 0; i < n; i++ {
		if ms[i] == '1' {
			ans++
		}
	}
	return ans
}

func diff(b byte) byte {
	if b == '1' {
		return '2'
	}
	return '1'
}

func count(b byte) int {
	if b == '1' {
		return 1
	}
	return 2
}
