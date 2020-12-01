package main

import (
	"fmt"
	"github.com/hashicorp/golang-lru"
)


func main(){
	l, e := lru.New(3)
	if e!=nil{
	}
	l.Add("1","2")
	l.Add("2","3")
	l.Add("3","4")
	l.Add("4","5")
	fmt.Println(l.Get("1"))
	fmt.Println(l.Get("4"))
}