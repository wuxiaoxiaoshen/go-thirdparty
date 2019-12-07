package main

import (
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func (K K8SClusterAction) GetCronJob(namespace string) {
	if namespace == "" {
		namespace = "default"
	}
	cronJobs,e := K.client.BatchV1beta1().CronJobs(namespace).List(v1.ListOptions{})
	if e!=nil{
		log.Println("k8s: cronJob fail")
		return
	}
	for _, i:=range cronJobs.Items{
		fmt.Println(i.Name, i.Status.String())
	}
}
