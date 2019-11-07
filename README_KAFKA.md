# Kafka

## 概念

- broker
- producer
- consumer
- topic
- partition
- offset

## 相关服务

- zookeeper
- kafka server
 
## 配置

- zookeeper
- kafka server

## 分区规则
- RoundRobin: 轮询
- Randomness: 随机
- Key: 按键
- 其他: 地理位置

## 压缩算法

- gzip
- zstd
- lz4
- snappy


## 客户端示例

> 单节点版本

- 节点信息
- topic 创建、删除
- 消息系统

**消息**

```golang
type Encoder interface {
	Encode() ([]byte, error)
	Length() int
}
```


**分区器**

```golang
type Partitioner interface {
	// Partition takes a message and partition count and chooses a partition
	Partition(message *ProducerMessage, numPartitions int32) (int32, error)

	// RequiresConsistency indicates to the user of the partitioner whether the
	// mapping of key->partition is consistent or not. Specifically, if a
	// partitioner requires consistency then it must be allowed to choose from all
	// partitions (even ones known to be unavailable), and its choice must be
	// respected by the caller. The obvious example is the HashPartitioner.
	RequiresConsistency() bool
}
```

**生产者**

**消费者**

**消费组**

- 重平衡机制： Sticky, RoundRobin, Range


> 集群版本

集群版本的表现形式：

本质：
- 多服务
- 多主机

服务多，主机多，直接如何通信，如何消息传递等。

对外：
- 单服务
- 单主机

只通过一个 broker，能访问到所有的 broker 信息。

创建一个 TOPIC，如果不备份，分区会分布在不同的 broker 上。
创建一个 TOPIC，如果备份数小于 broker 的数目，分区会分布在各 broker 上。比如 8 分区，2备份，那么broker 上共同存储 16个分区的日志
创建一个 TOPIC，如果备份数大于 broker 的数目，创建失败。

## 运维监控

- kafka-manager
- kafkacat