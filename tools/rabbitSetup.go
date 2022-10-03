package tools

import (
	"context"
	"github.com/vsnnkv/btcApplicationGo/config"
	"github.com/vsnnkv/btcApplicationGo/tools/rabbitmq"
	"log"
)

var rabbit rabbitmq.RabbitMQ

func Start(ctx context.Context) {
	rabbit = rabbitmq.NewRabbitMQ()

	setupRabbit(ctx)
}

func Shutdown(ctx context.Context) (done chan struct{}) {
	done = rabbit.Close(ctx)
	return
}

func setupRabbit(ctx context.Context) {
	var setup rabbitmq.Setup = func() {
		createQueues(rabbit)
	}
	cfg := config.Get()

	configConn := rabbitmq.ConfigConnection{
		URI:           cfg.RabbitUrl,
		PrefetchCount: 1,
	}
	rabbitmq.KeepConnectionAndSetup(ctx, rabbit, configConn, setup)
}

func createQueues(rabbit rabbitmq.QueueCreator) {
	config := rabbitmq.ConfigQueue{
		Name:       "logs",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
	_, err := rabbit.CreateQueue(config)
	if err != nil {
		log.Printf("error creating queue: %s\n", err)
	}
}

func Publish(ctx context.Context, msg string) {
	config := rabbitmq.ConfigPublish{
		Exchange:   "",
		RoutingKey: "logs",
	}
	if err := rabbit.Publish(ctx, []byte(msg), config); err != nil {
		log.Printf("error publishing message: %s\n", err)
	}
}
