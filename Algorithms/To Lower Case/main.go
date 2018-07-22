package main

import "fmt"

func main() {

	fmt.Println(toLowerCase("al&phaBET"))

}

func toLowerCase(str string) string {

	dest := []byte(str)
	for i, c := range dest {
		if c >= 'A' && c <= 'Z' {
			dest[i] = c + 32
		}else{
			continue
		}
	}
	return string(dest)
}
