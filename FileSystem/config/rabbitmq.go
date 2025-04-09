package config

const (
	// AsyncTransferEnable ：是否开启异步转移
	AsyncTransferEnable = true
	// RabbitURL ：rabbitmq服务的入口url
	RabbitURL = "amqp://admin:admin@49.232.28.223:5672/"
	// TransExchangeName ：用于文件transfer的交换金
	TransExchangeName = "uploadserver.trans"
	// TransCephQueueName ：ceph转移队列名
	TransCephQueueName = "uploadserver.trans.ceph"
	// TransCephErrQueueName ：ceph转移失败后写入另一个队列的队列名
	TransCephErrQueueName = "uploadserver.trans.ceph.err"
	// TransCephRoutingKey ：routingkey
	TransCephRoutingKey = "ceph"
)
