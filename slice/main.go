package main

import "fmt"

func main() {

	a := []int{1, 2, 3}
	fmt.Println(&a[0])
	b := append(a, 4, 5, 6)
	a = append(a, 4, 5, 6, 7, 8, 9)
	fmt.Println(a)
	fmt.Println(b)

	//fmt.Println(&a[0])
	//fmt.Println(len(a))
	//fmt.Println(cap(a))

	//var ss []int
	//ss = append(ss, 1, 2, 3)
	//arr := [10]int{1,2,3,4,5,6,7,8,9,0}
	//s1 := arr[1:5]
	//fmt.Println(cap(s1))

	//a1 := make([]int, 4, 8)
	//fmt.Println(&a1[0])
	//b1 := append(a1, 3, 3, 3)
	//a1 = append(a1, 4, 4, 4)
	//fmt.Println(&a1[0])
	//fmt.Println(&b1[0])
	//fmt.Println(a1)
	//fmt.Println(b1)

	var c []int64
	c = append(c, 1)
	fmt.Println(c)

}
