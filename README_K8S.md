# k8s

### 1. 安装部署

> 安装工具

- kubeadm
- kubectl
- kubelet

**kubeadm**

- preflight check

检查：image, cpu, memory, namespace, linux Controller Groups

- /etc/kubernetes

```text
admin.conf

```

**kubectl**

- 部署组件
- 查询节点信息


```text
kubectl cluster-info
kubectl get nodes
kubectl get deployments
kubectl describe node ${name}
kubectl get pods 
kubectl get pods -n kube-system
kubectl apply -f application.yaml
kubectl delete -f application.yaml


```


### 2. 容器编排

> 随心所欲部署自己的容器化应用

yaml: 规范

- 字段：驼峰式，首字母小写


公共字段：

```text
apiVersion
kind
metadata
spec
```



- Pod
- ReplicaSet
- Service
- Secret
- ConfigMap
- Deployment
- PersistentVolumeClaim
- PersistentVolume
- Job
- CronJob
- HorizontalPodAutoscaler
- Role/RoleBinding
- ClusterRole/ClusterRoleBinding


### 3. 各组建的交互逻辑

**master**

**node**
