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

