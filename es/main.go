package main

import "fmt"

func main(){
	//es := DefaultElasticClientAction
	//fmt.Println(es.IndexExists("users"))
	//fmt.Println(es.GetDoc("users", "1"))
	es := DefaultElasticSearchAction
	r:= es.client.Info
	fmt.Println(r)
}
