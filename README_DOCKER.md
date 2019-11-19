# DOCKER

## 初级阶段

- 镜像操作
- 容器操作
- 构建镜像：DockerFile

## 中级阶段

docker 实现原理：

- linux namespace: 隔离
- linux controller groups:  资源限制
- aufs(advance union file system)

需要关注的目录：

> /var/lib/docker/aufs
```text
diff  layers  mnt
```

- 文件挂载
- 层: 1. 只读层 2. init 层（/etc/hosts /etc/hostname) 3. 可读可写层

> /var/lib/docker

```text
aufs  containers  image  network  plugins  swarm  tmp  tmp-old  trust  volumes
```

- volume 默认挂载点：/var/lib/docker/volume/sha256/_data

> mount -t cgroup (/sys/fs/cgroup)

- 资源限制的文件系统，即创建子目录，写入pid, 值等能限制进行的资源消耗，比如 cpu, memory 等

> /proc/{pid}

- setns(): namespace 如何进入到一个进程中，即 docker exec 的实现机制(一个进程，可以选择加入到某个进程已有的 Namespace 当中，从而达到“进入”这个进程所在容器的目的)

> $HOME/.docker 

- docker 配置文件

## 中高级阶段

- DockerFile 构建进行，指令熟练
- Docker-compose 单节点上编排多个容器，以及容器之间的连接关系，指令熟练
- k8s 


