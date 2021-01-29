package main

import (
	"fmt"
	"reflect"
)

func main() {
	//author := "draven"
	//fmt.Println("TypeOf author:", reflect.TypeOf(author))
	//fmt.Println("ValueOf author:", reflect.ValueOf(author))
	//
	//year := 2020
	//fmt.Println("TypeOf year:", reflect.TypeOf(year))
	//fmt.Println("ValueOf year:", reflect.ValueOf(year))

	v := reflect.ValueOf(1)
	t := reflect.TypeOf(1)
	fmt.Println(t)
	fmt.Println(v)
	fmt.Println(t.Name())
	fmt.Println(t.Size())
	fmt.Println(v.Interface().(int))

}
