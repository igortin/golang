package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(func_check(10))
}

func countAndSay(n int) string {

	if n == 1 {
		return "1"
	}
	var tmp = "11"
	for i:=2 ; i < n ; i++ {
		tmp = compute(tmp)
	}
	return tmp

}


func compute(tmp string) string {
	var i = 0
	var temp = ""
	var flag = 1
	for j := 1; j < len(tmp); j++ {
		if tmp[i] == tmp[j] {
			i = j
			flag++
			if j + 1 == len(tmp){
				temp = temp +strconv.Itoa(flag) + string(tmp[i])
			}
		} else {
			temp = temp + strconv.Itoa(flag) + string(tmp[i])
			flag = 1
			i++
			if j + 1 == len(tmp){
				temp = temp +strconv.Itoa(flag) + string(tmp[i])
			}
		}
	}
	return temp
}