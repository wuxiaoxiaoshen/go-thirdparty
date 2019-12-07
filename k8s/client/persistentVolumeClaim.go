package main

import (
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func (K K8SClusterAction) GetPVC(namespace string) {
	if namespace == "" {
		namespace = "default"
	}
	pvcs, e := K.client.CoreV1().PersistentVolumeClaims(namespace).List(v1.ListOptions{})
	if e!=nil{
		log.Println("k8s: persistentVolumeClaims fail")
		return
	}
	for _, i:= range pvcs.Items{
		fmt.Println(i.Name, i.Spec.VolumeName)
	}

}
