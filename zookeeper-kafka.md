

-[kafka 安装](http://blog.csdn.net/hxpjava1/article/details/19159501)
-[官网quickstart](http://kafka.apache.org/documentation.html#quickstart)

## zookeeper

bin/zookeeper-server-start.sh config/zookeeper.properties

bin/zookeeper-server-start.sh config/zookeeper.properties &

“&”号是为了让在后台运行，要不还要在手动放后台或者重新开启一个终端窗口。

## kafka

**启动**
bin/kafka-server-start.sh config/server.properties &

kafka-server-start.bat ../../config/server.properties

**创建一个topic，名字叫test**
bin/kafka-topics.sh --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic test

kafka-topics.bat --create --zookeeper localhost:2181 --replication-factor 1 --partitions 1 --topic my-messages

**查看创建的topic**
bin/kafka-topics.sh --list --zookeeper localhost:2181 2>/dev/null
“2>/dev/null”屏蔽一些线程的启动信息等,只显示想看到的结果。

kafka-topics.bat --list --zookeeper localhost:2181

**kafka发送的message可以是文件或者标准输入的一行。**
bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test


kafka-console-producer.bat --broker-list localhost:9092 --topic my-messages

**收消息**
bin/kafka-console-consumer.sh --zookeeper localhost:2181 --topic test --from-beginning

kafka-console-consumer.bat --zookeeper localhost:2181 --topic my-messages --from-beginning

**多broker测试**
cp config/server.properties config/server-1.properties
cp config/server.properties config/server-2.properties

config/server-1.properties:
    broker.id=1
    port=9093
    log.dir=/tmp/kafka-logs-1
 
config/server-2.properties:
    broker.id=2
    port=9094
    log.dir=/tmp/kafka-logs-2

**启动broker1和broker2**

JMX_PORT=9997 bin/kafka-server-start.sh config/server-1.properties &
JMX_PORT=9998 bin/kafka-server-start.sh config/server-2.properties &

partition：同一个topic下可以设置多个partition，将topic下的message存储到不同的partition下，目的是为了提高并行性
leader：负责此partition的读写操作，每个broker都有可能成为某partition的leader
replicas：副本，即此partition在那几个broker上有备份，不管broker是否存活
isr：存活的replicas

## issue

Address family not supported by protocol family: connect
原因：
Zookeeper无法连接localhost(127.0.0.1)。

解决方案：

修改本机C盘hosts文件，将127.0.0.1映射成localhost。

C:\Windows\System32\drivers\etc

