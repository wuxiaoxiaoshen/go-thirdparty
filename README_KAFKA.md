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




> 集群版本

