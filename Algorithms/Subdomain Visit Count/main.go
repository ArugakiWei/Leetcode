package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}

func subdomainVisits(cpdomains []string) []string {

	sum := make(map[string]int)
	var result []string
	for _, str := range cpdomains {
		timeDomains := strings.Split(str, " ")
		time, _ := strconv.Atoi(timeDomains[0])
		for i := 1; i < len(strings.Split(timeDomains[1], "."))+1; i++ {
			cur := strings.SplitN(timeDomains[1], ".", i)
			last := len(cur) - 1
			sum[cur[last]] += time
		}
	}
	for key, val := range sum {
		str := fmt.Sprintf("%d %s", val, key)
		result = append(result, str)
	}
	return result
}
