package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type K8SClusterAction struct {
	client *kubernetes.Clientset
}

func NewK8SClusterAction(path string) *K8SClusterAction {
	home := os.Getenv("HOME")
	kubeConfig := filepath.Join(home, ".kube", "config")
	config, e := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if e != nil {
		log.Println("k8s: kubeConfig fail")
		return nil
	}
	client, e := kubernetes.NewForConfig(config)
	if e != nil {
		log.Println("k8s: client fail")
		return nil
	}
	return &K8SClusterAction{client: client}
}

var DefaultK8SClusterAction = &K8SClusterAction{}

func (K *K8SClusterAction) GetPods(namespace string) {
	pods, e := K.client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if e != nil {
		log.Println("k8s: get pods fail")
		return
	}
	fmt.Println("pods number", len(pods.Items))
	for _, i := range pods.Items {
		fmt.Println(i.Name)
	}

}
