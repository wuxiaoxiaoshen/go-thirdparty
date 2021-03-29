package main

import "fmt"

type A struct {
	Name string
}

func main(){

	var a *A
	var b *A
	b = new(A)
	if b == nil {
		fmt.Println("b", true)
	}
	if a == nil {
		fmt.Println("a", true)
	}
}