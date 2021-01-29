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

	//append()
}
