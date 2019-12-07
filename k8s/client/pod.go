package main

import (
	"fmt"
	 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

func (K *K8SClusterAction) GetPods(namespace string) []string{
	pods, e := K.client.CoreV1().Pods(namespace).List(v1.ListOptions{})
	if e != nil {
		log.Println("k8s: get pods fail")
		return nil
	}
	var result []string
	fmt.Println(fmt.Sprintf("namespace: %s, pods number: %d", namespace, len(pods.Items)))
	for _, i := range pods.Items {
		fmt.Println("Pod Name:", i.Name)
		result = append(result, i.Name)
	}
	return result

}
