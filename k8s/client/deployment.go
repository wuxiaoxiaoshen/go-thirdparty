package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func (K *K8SClusterAction) GetDeployment(namespace string) {
	if namespace == "" {
		namespace =  v1.NamespaceDefault
	}
	deployment := K.client.AppsV1().Deployments(namespace)
	d,e := deployment.List(metav1.ListOptions{})
	if e!=nil{
		log.Println("k8s: deployment is nil")
		return
	}
	for _, i:=range d.Items{
		fmt.Println("Deployment Name: ",i.Name)
	}
}
