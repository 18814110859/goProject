package main

import (
	"fmt"
)

func main() {
	hash2 := make(map[string]int64, 26)
	vstatk := []string{"1", "2", "3", "4", "5", "6"}
	vstatv := []int64{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(vstatk); i++ {
		hash2[vstatk[i]] = vstatv[i]
	}

	fmt.Println(hash2)
	v, ok := hash2["1"]
	fmt.Println(v)
	fmt.Println(ok)

	//fmt.Println(len(hash2))
	//for k, v := range hash2 {
	//	fmt.Printf("%s = %d./n", k, v)
	//}
	//
	//delete(hash2, "1")// 删除k 是 "1"

	//hash1 := map[int]int{
	//	1:2,
	//	3:4,
	//	5:6,
	//}
	//fmt.Println(hash1)
	//
	//slice1 := make([]int,4, 8)
	//fmt.Println(slice1)
	//
	//slice2 := make([]string,8, 8)
	//fmt.Println(slice2)
	//
}
