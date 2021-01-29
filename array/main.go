package main

import (
	"fmt"
)

func main() {
	a := [200]interface{}{}
	fmt.Println(&a[0], &a[1])

	b := newArray()
	fmt.Println(b)

}

//go:noinline
func newArray() *[4]int {
	a := [4]int{1, 2, 3, 4}
	return &a
}
