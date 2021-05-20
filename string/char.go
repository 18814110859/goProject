package main

import "fmt"

func main() {
	str := "Hello,世界"
	var ch byte
	for i := 0; i < len(str); i++ {
		ch = str[i]
		fmt.Println(i, ch)
	}

	for i, char := range str {
		fmt.Println(i, char)
	}


	var flag1 = false
	var flag2 = false
	flag1 = true
	flag2 = true
	
}
