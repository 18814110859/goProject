package main

import "fmt"

func main() {
	_, _, nikeName := GetName()
	fmt.Println(nikeName)
}

func GetName() (firstName string, lastName string, nickName string) {
	return "May", "Chan", "Chibi Maruko"
}
