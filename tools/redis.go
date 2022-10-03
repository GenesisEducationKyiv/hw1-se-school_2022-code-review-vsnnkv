package tools

import (
	"context"
	"github.com/vsnnkv/btcApplicationGo/tools/rabbitmq"
	"log"
	"os"
)

var rabbit rabbitmq.RabbitMQ

// Start starts the RabbitMQ connection
func Start(ctx context.Context) {
	rabbit = rabbitmq.NewRabbitMQ()

	setupRabbit(ctx)
}

// Shutdown stops the RabbitMQ connection
func Shutdown(ctx context.Context) (done chan struct{}) {
	done = rabbit.Close(ctx)
	return
}

func setupRabbit(ctx context.Context) {
	var setup rabbitmq.Setup = func() {
		createQueues(rabbit)
	}
	configConn := rabbitmq.ConfigConnection{
		URI:           "amqp://guest:guest@localhost:5672",
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

// PublishTest publishes a test message to the RabbitMQ exchange
func PublishTest(ctx context.Context, msg string) {
	config := rabbitmq.ConfigPublish{
		Exchange:   "",
		RoutingKey: "logs",
	}
	if err := rabbit.Publish(ctx, []byte(msg), config); err != nil {
		log.Printf("error publishing message: %s\n", err)
	}
}

func loadURI() (uri string) {
	uri = os.Getenv("RABBITMQ_URI")
	return
}

//var wg *sync.WaitGroup = &sync.WaitGroup{}
//
//type rabbit struct {
//	conn       *amqp.Connection
//	chConsumer *amqp.Channel
//	chProducer *amqp.Channel
//	wg         *sync.WaitGroup
//}
//type ConfigPublish struct {
//	Exchange        string
//	RoutingKey      string
//	Mandatory       bool
//	Immediate       bool
//	Headers         amqp.Table
//	ContentType     string
//	ContentEncoding string
//	Priority        uint8
//	CorrelationID   string
//	MessageID       string
//}
//
//// NewRabbitMQ creates the object to manage the operations to rabbitMQ
//func NewRabbitMQ() *rabbit {
//	return &rabbit{
//		wg: &sync.WaitGroup{},
//	}
//}
//
//func StartRabbit(ctx context.Context, prefetchCount int) {
//	go func() {
//		for {
//			notifyClose, err := connect(prefetchCount)
//			if err != nil {
//				log.Printf("error connecting to rabbitmq: [%s]\n", err)
//				time.Sleep(time.Second * 5)
//				continue
//			}
//			// create queues, exchanges, consume from queue, etc
//			select {
//			case <-notifyClose:
//				continue
//			case <-ctx.Done():
//				return
//			}
//		}
//	}()
//}
//
//func connect(prefetchCount int) (notify chan *amqp.Error, err error) {
//	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
//	if err != nil {
//		return
//	}
//	chConsumer, err := conn.Channel()
//	if err != nil {
//		return
//	}
//	if prefetchCount > 0 {
//		err = chConsumer.Qos(prefetchCount, 0, false)
//		if err != nil {
//			return
//		}
//	}
//	notify = make(chan *amqp.Error)
//	conn.NotifyClose(notify)
//	return
//}
//
//func (r *rabbit) Publish(ctx context.Context, body []byte, config ConfigPublish) (err error) {
//	if r.chConsumer == nil {
//		return amqp.ErrClosed
//	}
//	r.wg.Add(1)
//	defer r.wg.Done()
//	err = r.chProducer.PublishWithContext(
//		ctx,
//		config.Exchange,
//		config.RoutingKey,
//		config.Mandatory,
//		config.Immediate,
//		amqp.Publishing{
//			Headers:         config.Headers,
//			ContentType:     config.ContentType,
//			ContentEncoding: config.ContentEncoding,
//			Priority:        config.Priority,
//			CorrelationId:   config.CorrelationID,
//			MessageId:       config.MessageID,
//			Body:            body,
//		},
//	)
//	return
//}
//
//var rabbit1 rabbit
//
//func PublishTest(ctx context.Context, msg string) {
//	config := ConfigPublish{
//		Exchange:        "",
//		RoutingKey:      "test",
//		Mandatory:       false,
//		Immediate:       false,
//		Headers:         nil,
//		ContentType:     "",
//		ContentEncoding: "utf-8",
//		Priority:        0,
//		CorrelationID:   "",
//		MessageID:       "",
//	}
//
//	if err := rabbit1.Publish(ctx, []byte(msg), config); err != nil {
//		log.Printf("error publishing message: %s\n", err)
//	}
//
//}

//type Publisher interface {
//	Publish(exchange, RoutingKey string, body []byte)
//}

//type publisher struct{}
//
//func (p publisher) Publish(exchange, RoutingKey string, body []byte) {
//	rabbit.publish(exchange, RoutingKey, body)
//}

//func failOnError(err error, msg string) {
//	if err != nil {
//		log.Panicf("%s: %s", msg, err)
//	}
//}
//
//func SomeFuncRabbit() {
//
//	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
//	failOnError(err, "Failed to connect to RabbitMQ")
//	defer conn.Close()
//
//	ch, err := conn.Channel()
//	failOnError(err, "Failed to open a channel")
//	defer ch.Close()
//
//	err = ch.ExchangeDeclare(
//		"logs",   // name
//		"fanout", // type
//		true,     // durable
//		false,    // auto-deleted
//		false,    // internal
//		false,    // no-wait
//		nil,      // arguments
//	)
//	failOnError(err, "Failed to declare an exchange")
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	body := bodyFrom(os.Args)
//	err = ch.PublishWithContext(ctx,
//		"logs", // exchange
//		"",     // routing key
//		false,  // mandatory
//		false,  // immediate
//		amqp.Publishing{
//			ContentType: "text/plain",
//			Body:        []byte(body),
//		})
//	failOnError(err, "Failed to publish a message")
//
//	log.Printf(" [x] Sent %s", body)
//}
//
//func bodyFrom(args []string) string {
//	var s string
//	if (len(args) < 2) || os.Args[1] == "" {
//		s = "hello"
//	} else {
//		s = strings.Join(args[1:], " ")
//	}
//	return s
//}
