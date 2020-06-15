## kafka

```$xslt
Kafka是一个分布式的、分区化、可复制提交的日志服务
Kafka实现了不同应用程序之间的松耦和，那么作为一个可扩展、高可靠的消息系统
支持高Throughput(吞吐量)的应用
scale out：无需停机即可扩展机器
持久化：通过将数据持久化到硬盘以及replication防止数据丢失
支持online和offline的场景
```

#### Kafka的特点
```$xslt
Kafka是分布式的，其所有的构件borker(服务端集群)、producer(消息生产)、consumer(消息消费者)都可以是分布式的
在消息的生产时可以使用一个标识topic来区分，且可以进行分区；每一个分区都是一个顺序的、不可变的消息队列， 并且可以持续的添加。

同时为发布和订阅提供高吞吐量。据了解，Kafka每秒可以生产约25万消息（50 MB），每秒处理55万消息（110 MB）。

消息被处理的状态是在consumer端维护，而不是由server端维护。当失败时能自动平衡

```

#### Kafka选择分区的模式（3种）
     
```$xslt
     指定往哪个分区写
     指定key，kafka根据key做hash然后决定写哪个分区
     轮询方式
```

#### 生产者往kafka发送数据的模式（3种
```$xslt
0：把数据发给leader就成功，效率最高、安全性最低。
1: 把数据发送给leader，等待leader回ACK
all:把数据发给leader,确保follower从leader拉取数据回复ack给leader，leader再回复ACK；安全性最高

```

#### 常用的场景
```$xslt
1. 监控：主机通过Kafka发送与系统和应用程序健康相关的指标，然后这些信息会被收集和处理从而创建监控仪表盘并发送警告。
2. 消息队列： 应用程度使用Kafka作为传统的消息系统实现标准的队列和消息的发布—订阅,
比起大多数的消息系统来说，Kafka有更好的吞吐量，内置的分区，冗余及容错性，这让Kafka成为了一个很好的大规模消息处理应用的解决方案。消息系统 一般吞吐量相对较低，但是需要更小的端到端延时，并尝尝依赖于Kafka提供的强大的持久性保障。
3. 站点的用户活动追踪: 为了更好地理解用户行为，改善用户体验，将用户查看了哪个页面、点击了哪些内容等信息发送到每个数据中心的Kafka集群上，并通过Hadoop进行分析、生成日常报告。
4. 流处理：保存收集流数据，以提供之后对接的Storm或其他流式计算框架进行处理。
5. 日志聚合:使用Kafka代替日志聚合（log aggregation）。日志聚合一般来说是从服务器上收集日志文件，然后放到一个集中的位置（文件服务器或HDFS）进行处理。
6. 持久性日志：Kafka可以为一种外部的持久性日志的分布式系统提供服务。这种日志可以在节点间备份数据，并为故障节点数据回复提供一种重新同步的机制。
```

#### Kafka中包含以下基础概念
```$xslt
 1.Topic(话题)：Kafka中用于区分不同类别信息的类别名称。由producer指定
    2.Producer(生产者)：将消息发布到Kafka特定的Topic的对象(过程)
    3.Consumers(消费者)：订阅并处理特定的Topic中的消息的对象(过程)
    4.Broker(Kafka服务集群)：已发布的消息保存在一组服务器中，称之为Kafka集群。集群中的每一个服务器都是一个代理(Broker). 消费者可以订阅一个或多个话题，并从Broker拉数据，从而消费这些已发布的消息。
    5.Partition(分区)：Topic物理上的分组，一个topic可以分为多个partition，每个partition是一个有序的队列。partition中的每条消息都会被分配一个有序的id（offset）
    Message：消息，是通信的基本单位，每个producer可以向一个topic（主题）发布一些消息。
```

#### ⼯作流程
```$xslt
 1.⽣产者从Kafka集群获取分区leader信息
    2.⽣产者将消息发送给leader
    3.leader将消息写入本地磁盘
    4.follower从leader拉取消息数据
    5.follower将消息写入本地磁盘后向leader发送ACK
    6.leader收到所有的follower的ACK之后向生产者发送ACK
```

#### Topic和数据⽇志
     topic 是同⼀类别的消息记录（record）的集合
     每个partition都是⼀个有序并且不可变的消息记录集合。
     Kafka可以配置⼀个保留期限，⽤来标识⽇志会在Kafka集群内保留多⻓时间。Kafka集群会保留在保留 期限内所有被发布的消息，不管这些消息是否被消费过。
     
#### Partition结构
     Partition在服务器上的表现形式就是⼀个⼀个的⽂件夹，每个partition的⽂件夹下⾯会有多组segment ⽂件，每组segment⽂件⼜包含 .index ⽂件、 .log ⽂件、 .timeindex ⽂件三个⽂件，其中 .log ⽂ 件就是实际存储message的地⽅，⽽ .index 和 .timeindex ⽂件为索引⽂件，⽤于检索消息。
     
     