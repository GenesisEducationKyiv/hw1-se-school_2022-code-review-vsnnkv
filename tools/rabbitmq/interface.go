package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ interface {
	Connector
	Closer
	QueueCreator
	Publisher
}

type Connector interface {
	Connect(config ConfigConnection) (notify chan *amqp.Error, err error)
}

type Closer interface {
	Close(ctx context.Context) (done chan struct{})
}

type QueueCreator interface {
	CreateQueue(config ConfigQueue) (queue amqp.Queue, err error)
}

type Publisher interface {
	Publish(ctx context.Context, body []byte, config ConfigPublish) (err error)
}

type RabbitSetup interface {
	Setup()
}

type Setup func()

func (s Setup) Setup() {
	s()
}
