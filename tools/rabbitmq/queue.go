package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func (r *rabbit) CreateQueue(config ConfigQueue) (queue amqp.Queue, err error) {
	if r.chConsumer == nil {
		err = amqp.ErrClosed
		return
	}
	queue, err = r.chConsumer.QueueDeclare(
		config.Name,
		config.Durable,
		config.AutoDelete,
		config.Exclusive,
		config.NoWait,
		config.Args,
	)
	return
}
