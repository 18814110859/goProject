package main

import "fmt"

type MyStruct struct {
	i int
	j int
}

// 传递struct指针时，会拷贝struct指针 然后分配一个新的内存 来 存储这个指针
func myFunction(b *MyStruct) {
	b.i = 41
	b.j = 31
	fmt.Printf("in my_function - b=(%v, %p, %p)\n", b, b, &b)
}

func main() {
	b := &MyStruct{i: 40, j: 30}
	fmt.Printf("before calling - b=(%v, %p, %p)\n", b, b, &b)
	myFunction(b)
	fmt.Printf("after calling  - b=(%v, %p, %p)\n", b, b, &b)
}
