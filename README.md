<h1 align="center">Go-ThirdParty</h1>
<p align="center">
    <a href="https://github.com/wuxiaoxiaoshen">
        <img src="https://img.shields.io/badge/Author-wuxiaoxiaoshen-green" alt="Author">
        <img src="https://img.shields.io/badge/Project-GoThirdParty-red" alt="Project">
    </a>
</p>

# 优秀第三方库使用指北

> go third package




## 命名规则

- 主题: Topic
- 代码: Code
- 模式: Action: Do: String

## 方法论

- 原理
- 使用

## 学习什么

- 代码组织
- 命名方式
- 代码方式

## kafka
> tips: cd kafka && go run *.go

- [kafka 说明](README_KAFKA.md)

## NSQ

- [NSQ 说明](README_NSQ.md)

## Redis

## ElasticSearch

## Kubernetes

## Docker

## Mysql



## 抓包

> 分析网络、client 和 Server 之间的交互协议

工具：tcpdump, wireshark

- tcpdump 抓取完整的包
- wireshark 对应协议查看包的传输过程

最常用法：tcp 三次握手、四次挥手

Example: REDIS.sh(REDIS.cap), MONGO.sh(MONGO.cap)