package config

import "os"

const (
	// AsyncTransferEnable ：是否开启异步转移
	AsyncTransferEnable = true
	// TransExchangeName ：用于文件transfer的交换金
	TransExchangeName = "uploadserver.trans"
	// TransCephQueueName ：ceph转移队列名
	TransCephQueueName = "uploadserver.trans.ceph"
	// TransCephErrQueueName ：ceph转移失败后写入另一个队列的队列名
	TransCephErrQueueName = "uploadserver.trans.ceph.err"
	// TransCephRoutingKey ：routingkey
	TransCephRoutingKey = "ceph"
)

var (
	// RabbitURL ：rabbitmq服务的入口url
	RabbitURL = os.Getenv("RABBIT_URL")
)
