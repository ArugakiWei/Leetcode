package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compareVersion("1.01", "1.001"))
	fmt.Println(compareVersion("1.0", "1.0.0"))
	fmt.Println(compareVersion("0.1", "1.1"))
	fmt.Println(compareVersion("1.1", "1.10"))
	fmt.Println(compareVersion("1.2", "1.10"))
	fmt.Println(compareVersion(
		"19.8.3.17.5.01.0.0.4.0.0.0.0.0.0.0.0.0.0.0.0.0.00.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.000000.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.000000",
		"19.8.3.17.5.01.0.0.4.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0000.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.0.000000"))
}

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	for i, v := range v1 {
		if isAllZero(v) {
			v1[i] = "0"
		} else {
			v1[i] = strings.TrimLeft(v, "0")
		}
	}
	for i, v := range v2 {
		if isAllZero(v) {
			v2[i] = "0"
		} else {
			v2[i] = strings.TrimLeft(v, "0")
		}
	}

	minLen := len(v1)
	if len(v2) < len(v1) {
		minLen = len(v2)
	}
	var i int
	for i = 0; i < minLen; i++ {
		vv1, _ := strconv.Atoi(v1[i])
		vv2, _ := strconv.Atoi(v2[i])
		if vv1 > vv2 {
			return 1
		}
		if vv1 < vv2 {
			return -1
		}
	}
	if len(v1) > len(v2) {
		for ; i < len(v1); i++ {
			if v1[i] != "0" {
				return 1
			}
		}
	}
	if len(v2) > len(v1) {
		for ; i < len(v2); i++ {
			if v2[i] != "0" {
				return -1
			}
		}
	}
	return 0
}

func isAllZero(s string) bool {
	for _, v := range s {
		if v != '0' {
			return false
		}
	}
	return true
}
