package main

import "fmt"

func main() {
	// 创建一个数组
	a := [3]int{1, 2, 3}
	fmt.Printf("before: a=(%v, %p)\n", a, &a)
	fmt.Printf("before: a[0]=(%v, %p)\n", a[0], &a[0])
	fmt.Printf("before: a[1]=(%v, %p)\n", a[1], &a[1])
	changeA(a[0:2])
	fmt.Printf("after: a=(%v, %p)\n", a, &a)
	fmt.Printf("after: a[0]=(%v, %p)\n", a[0], &a[0])
	fmt.Printf("after: a[1]=(%v, %p)\n", a[1], &a[1])
}

func changeA(a []int) {
	a[0] = 100
	fmt.Printf("inside: a=(%v, %p)\n", a, &a)
	fmt.Printf("inside: a[0]=(%v, %p)\n", a[0], &a[0])
	fmt.Printf("inside: a[1]=(%v, %p)\n", a[1], &a[1])
}
