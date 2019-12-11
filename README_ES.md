# ElasticSearch


### 单节点
> es

```text
curl http://127.0.0.1:9200
curl http://127.0.0.1:9200/_cat/health
curl http://127.0.0.1:9200/_cat/nodes
```

### 多节点: 集群

> es_01, es_02, es_03

```text
curl http://127.0.0.1:9200
curl http://127.0.0.1:9200/_cat/health
curl http://127.0.0.1:9200/_cat/nodes
```

### ELK

> es(搜索), logstash(日志采集), kibana(可视化), cerebro(监控)

### 使用清单

#### 概念

- 集群
- 节点
- 分片
- 索引
- 类型
- 文档
- 字段

### 清单

- 索引 index

1。自定义 mapping， 对字段进行类型和是否索引设置，层级，字段最大个数，字段最大长度等。
```text
text
integer
float
long
date
boolean
ip
```
2。dynamic mapping 不显式的自定义 mapping, 则使用默认的自定义的 mapping 规则

```text
POST /index
{
   "mapping": {
       "properties" : {
            "filed": { "type": "boolean"}
        }
   }
}

PUT /index/_mapping
{
    "properties": {
        "filed_new": {"type": "boolean"}
    }
}

GET /index/_mapping

GET /index/_mapping/filed/name,age
```

- curd document

1。获取


