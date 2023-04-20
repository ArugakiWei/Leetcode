package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countDaysTogether("08-15", "08-18", "08-16", "08-19"))
	fmt.Println(countDaysTogether("10-01", "10-31", "11-01", "12-31"))
	fmt.Println(countDaysTogether("09-01", "10-19", "06-19", "10-20"))
	fmt.Println(countDaysTogether("08-06", "12-08", "02-04", "09-01"))
}

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	aa, la := to365(arriveAlice), to365(leaveAlice)
	ab, lb := to365(arriveBob), to365(leaveBob)
	fmt.Println(aa, ab, la, lb)
	if aa <= ab {
		if ab > la {
			return 0
		}
		if ab <= la && lb >= la {
			return la - ab + 1
		}
		return lb - ab + 1
	}
	if aa > lb {
		return 0
	}
	if aa <= lb && la >= lb {
		return lb - aa + 1
	}
	return la - aa + 1
}

var Days = []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}

func to365(s string) int {
	m := s[:2]
	d := s[3:5]
	a, _ := strconv.Atoi(m)
	if m[0] == '0' {
		a, _ = strconv.Atoi(string(m[1]))
	}
	b, _ := strconv.Atoi(d)
	if d[0] == '0' {
		b, _ = strconv.Atoi(string(d[1]))
	}
	return Days[a-1] + b
}
