package main

import (
	"fmt"
	"github.com/hashicorp/golang-lru"
	"runtime"
)


func main(){
	l, e := lru.New(10000000000)
	if e!=nil{
	}
	//for i:=0;i<1000000;i++{
	//	l.Add(fmt.Sprintf("%d",i), fmt.Sprintf("%d",i))
	//}
	fmt.Println(l.Get("1"))
	fmt.Println(l.Get("4"))
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println(m.Sys/1024/1024)
}