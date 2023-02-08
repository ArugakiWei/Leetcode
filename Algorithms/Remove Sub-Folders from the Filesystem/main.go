package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeSubfolders([]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}))
	fmt.Println(removeSubfolders([]string{"/a", "/a/b/c", "/a/b/d"}))
	fmt.Println(removeSubfolders([]string{"/a/b/c", "/a/b/ca", "/a/b/d"}))
}

func removeSubfolders(folder []string) []string {
	m := make(map[string]struct{})
	for _, p := range folder {
		m[p] = struct{}{}
	}
	for _, p := range folder {
		prefix := ""
		fs := strings.Split(p, "/")
		for i := 0; i < len(fs); i++ {
			if fs[i] == "" {
				continue
			}
			prefix += "/" + fs[i]
			if prefix == p {
				break
			}
			if _, ok := m[prefix]; ok {
				delete(m, p)
			}
		}
	}

	var ans []string
	for v := range m {
		ans = append(ans, v)
	}
	return ans
}
