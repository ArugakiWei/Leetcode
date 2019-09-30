package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}