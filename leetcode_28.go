package main

import "fmt"

func main() {

	fmt.Println(strStr("aabaaabaaac", "aabaaac"))
}
func strStr(haystack string, needle string) int {
	i := 0
	count := 0
	if len(needle) == 0 { return 0 }
	if len(haystack) == 0  { return -1 }
	for j := 0; j < len(haystack); j++ {
		if haystack[j] == needle[i] {
			i++
			count= count+1
			if count == len(needle){ return j - len(needle) + 1}
		} else {
			count = 0
			if j >= i { j = j - i }
			i = 0
		}
	}
	return -1
}
