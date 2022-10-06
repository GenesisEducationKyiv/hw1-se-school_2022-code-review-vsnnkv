package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// ConfigConnection is the configuration for the connection
type ConfigConnection struct {
	URI           string
	PrefetchCount int
}

// ConfigQueue is the configuration for the queue
type ConfigQueue struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

// ConfigPublish is the configuration for the publisher
type ConfigPublish struct {
	Exchange        string
	RoutingKey      string
	Mandatory       bool
	Immediate       bool
	Headers         amqp.Table
	ContentType     string
	ContentEncoding string
	Priority        uint8
	CorrelationID   string
	MessageID       string
}
