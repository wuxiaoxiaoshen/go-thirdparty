package main

import (
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func (K K8SClusterAction) GetPV(namespace string) {
	if namespace == "" {
		namespace = "default"
	}
	pvs,e := K.client.CoreV1().PersistentVolumeClaims(namespace).List(v1.ListOptions{})
	if e!=nil{
		log.Println("k8s: persistentVolume fail")
		return
	}
	for _, i:=range pvs.Items {
		fmt.Println(i.Name,i.Status, i.Spec.VolumeName)
	}
}
