package main

import (
	"encoding/json"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"log"
)
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func (K K8SClusterAction) GetConfigMap(namespace string) {
	if namespace == "" {
		namespace = v1.NamespaceDefault
	}
	configMaps, e := K.client.CoreV1().ConfigMaps(namespace).List(metav1.ListOptions{})
	if e!=nil{
		log.Println("k8s: configMaps fail")
		return
	}
	for _, i:= range configMaps.Items {
		fmt.Println("configMap name:",i.Name)
		result, _ := json.MarshalIndent(i.Data, " ", "")
		fmt.Println(string(result))
	}

}
