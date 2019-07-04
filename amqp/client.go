package amqp

import (
	"github.com/streadway/amqp"
	"sync"
)

// RabbitMQ 用于管理和维护rabbitmq的对象
type RabbitMQ struct {
	wg sync.WaitGroup

	channel      *amqp.Channel
	exchangeName string // exchange的名称
	exchangeType string // exchange的类型
	receivers    []Receiver
}

// New 创建一个新的操作RabbitMQ的对象
func New() *RabbitMQ {
	// 这里可以根据自己的需要去定义
	return &RabbitMQ{
		exchangeName: ExchangeName,
		exchangeType: ExchangeType,
	}
}
