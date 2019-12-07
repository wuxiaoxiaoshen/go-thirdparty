package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func (K K8SClusterAction) GetServices(namespace string) {
	if namespace == "" {
		namespace = v1.NamespaceDefault
	}
	services,e := K.client.CoreV1().Services(namespace).List(metav1.ListOptions{})
	if e!=nil{
		log.Println("k8s: services fail")
		return
	}
	for _, i:= range services.Items {
		fmt.Println(i.Name, )
	}
}
