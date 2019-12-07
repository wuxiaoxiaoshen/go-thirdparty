package main


var k8s *K8SClusterAction

func init() {
	k8s = NewK8SClusterAction("")
}

func main() {
	//Namespace()
	//ConfigMap()
	//Services()
	//PersistentVolume()
	//PersistentVolumeClaims()
	//Deployment()
	//Pod()
	Job()
	CronJob()
}

func Namespace() {
	k8s.GetNamespace()
}
func ConfigMap() {
	//k8s.GetConfigMap("default")
	k8s.GetConfigMap("xw-example")
}
func Services() {
	//k8s.GetServices("default")
	k8s.GetServices("xw-example")
}

func PersistentVolume() {
	k8s.GetPV("xw-example")
}

func PersistentVolumeClaims() {
	k8s.GetPVC("xw-example")
}

func Pod() {
	//k8s.GetPods("default")
	k8s.GetPods("xw-example")
}

func Deployment() {
	k8s.GetDeployment("default")
	k8s.GetDeployment("xw-example")
}
func Job(){
	k8s.GetJob("xw-example")
}

func CronJob() {
	k8s.GetCronJob("xw-example")
}