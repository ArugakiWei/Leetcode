package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(alertNames([]string{"c", "c", "c", "c", "c", "c", "c", "c"}, []string{"02:48", "09:11", "13:08", "17:05", "14:06", "16:01", "16:19", "10:50"}))
}

func alertNames(keyName []string, keyTime []string) []string {
	m := make(map[string][]int)
	for i, name := range keyName {
		m[name] = append(m[name], toMinute(keyTime[i]))
	}
	var ans []string
	for name, time := range m {
		if len(time) < 3 {
			continue
		}
		sort.Ints(time)
		count, i, j := 1, 0, 1
		for i < len(time) && j < len(time) {
			if time[j]-time[i] <= 60 {
				count++
				j++
				if count >= 3 {
					ans = append(ans, name)
					break
				}
				continue
			} else {
				count = 1
				i++
				j = i + 1
			}
		}
	}
	sort.Strings(ans)
	return ans
}

func toMinute(time string) int {
	h := time[:2]
	if time[0] == '0' {
		h = time[1:2]
	}
	m := time[3:5]
	if time[3] == '0' {
		m = time[4:5]
	}
	hour, _ := strconv.Atoi(h)
	min, _ := strconv.Atoi(m)
	return hour*60 + min
}
