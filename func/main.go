package main

func myFunction() int {
	var a int
	defer func() {
		a += 1
	}()
	return a
}

func main() {
	myFunction()

	//a := &MyStruct{i: 40,j: 50}
	//MyFunc1(a)
	//fmt.Printf("[%p] %v\n", a, a)
}

//type MyStruct struct {
//	i int
//	j int
//}
//
//func MyFunc1(ms *MyStruct)  {
//	ms.i = 1
//	ms.j = 1
//	fmt.Printf("[%p] %v\n", ms, ms)
//}
//
//func MyFunc(ms *MyStruct)  {
//	ptr := unsafe.Pointer(ms)
//	for i := 0; i < 2; i++ {
//		c := (*int)(unsafe.Pointer(uintptr(ptr) + uintptr(8*i)))
//		*c += i + 1
//		fmt.Println(c)
//		fmt.Printf("[%p] %d\n", c, *c)
//	}
//}
