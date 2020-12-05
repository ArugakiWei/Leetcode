package main

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertNum := 0
	for x > revertNum {
		revertNum = revertNum*10 + (x % 10)
		x = x / 10
	}

	return revertNum == x || x == revertNum/10
}
