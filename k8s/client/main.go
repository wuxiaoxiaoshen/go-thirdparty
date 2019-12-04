package main

func main() {
	k8s := NewK8SClusterAction("")
	k8s.GetPods("xw-example")
	k8s.GetDeployment("xw-example")
}
