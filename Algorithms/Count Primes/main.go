package main

import "fmt"

func main() {
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {
	nums := make([]bool, n)
	for i := 2; i < n; i++ {
		nums[i] = true
	}

	for i := 2; i*i < n; i++ {
		if isPrime(i) {
			for j := i * i; j < n; j += i {
				nums[j] = false
			}
		}
	}

	count := 0
	for i := 2; i < n; i++ {
		if nums[i] {
			count++
		}
	}
	return count
}

func isPrime(n int) bool {
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
