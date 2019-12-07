package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)
func (K K8SClusterAction) GetNamespace() {
	namespace,e := K.client.CoreV1().Namespaces().List(v1.ListOptions{})
	if e!=nil{
		log.Println("k8s: namespace fail")
		return
	}
	fmt.Println(fmt.Sprintf("namespace length: %d", len(namespace.Items)))
	for _,i :=range namespace.Items {
		fmt.Println(i.Name, i.Status)
	}
}
