package main

import (
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func (K K8SClusterAction) GetJob(namespace string) {
	if namespace == "" {
		namespace = "default"
	}
	jobs,e := K.client.BatchV1().Jobs(namespace).List(v1.ListOptions{})
	if e!=nil{
		log.Println("k8s: jobs fail")
		return
	}
	for _, i:= range jobs.Items {
		fmt.Println(i.Name, i.Status.String())
	}
}
