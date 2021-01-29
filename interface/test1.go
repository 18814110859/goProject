package main

import "fmt"

type Cat struct{}
type Duck interface {
	Quack()
}

func main() {
	var c Duck = &Cat{}
	c.Quack()
}

func (c *Cat) Quack() {
	fmt.Println("meow")
}
