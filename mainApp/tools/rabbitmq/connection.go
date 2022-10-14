package rabbitmq

import (
	"context"
	"log"
	"sync"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var notifyOpenConn, notifySetupDone []chan struct{}
var muxNotifyOpenConn, muxNotifySetup sync.Mutex = sync.Mutex{}, sync.Mutex{}

func (r *rabbit) Connect(config ConfigConnection) (notify chan *amqp.Error, err error) {
	r.conn, err = amqp.Dial(config.URI)
	if err != nil {
		return
	}
	r.chProducer, err = r.conn.Channel()
	if err != nil {
		return
	}
	r.chConsumer, err = r.conn.Channel()
	if err != nil {
		return
	}
	if config.PrefetchCount > 0 {
		err = r.chConsumer.Qos(config.PrefetchCount, 0, false)
		if err != nil {
			return
		}
	}
	notifyOpenConnections()
	notify = make(chan *amqp.Error)
	r.conn.NotifyClose(notify)
	return
}

func (r *rabbit) Close(ctx context.Context) (done chan struct{}) {
	done = make(chan struct{})

	doneWaiting := make(chan struct{})
	go func() {
		r.wg.Wait()
		close(doneWaiting)
	}()

	go func() {
		defer close(done)
		select {
		case <-doneWaiting:
		case <-ctx.Done():
		}
		closeConnections(r)
	}()
	return
}

func KeepConnectionAndSetup(ctx context.Context, conn Connector, config ConfigConnection, setupRabbit RabbitSetup) {
	go func() {
		for {
			notifyClose, err := conn.Connect(config)
			if err != nil {
				log.Printf("error connecting to rabbitmq: [%s]\n", err)
				time.Sleep(time.Second * 5)
				continue
			}
			setupRabbit.Setup()
			notifySetupIsDone()
			select {
			case <-notifyClose:
				continue
			case <-ctx.Done():
				return
			}
		}
	}()
}

func notifyOpenConnections() {
	muxNotifyOpenConn.Lock()
	defer muxNotifyOpenConn.Unlock()
	for _, notify := range notifyOpenConn {
		close(notify)
	}
	notifyOpenConn = make([]chan struct{}, 0)
}

func notifySetupIsDone() {
	muxNotifySetup.Lock()
	defer muxNotifySetup.Unlock()
	for _, notify := range notifySetupDone {
		close(notify)
	}
	notifySetupDone = make([]chan struct{}, 0)
}

func closeConnections(r *rabbit) {
	var err error
	if r.chConsumer != nil {
		err = r.chConsumer.Close()
		if err != nil {
			log.Printf("Error closing consumer channel: [%s]\n", err)
		}
	}
	if r.chProducer != nil {
		err = r.chProducer.Close()
		if err != nil {
			log.Printf("Error closing producer channel: [%s]\n", err)
		}
	}
	if r.conn != nil {
		err = r.conn.Close()
		if err != nil {
			log.Printf("Error closing connection: [%s]\n", err)
		}
	}
}
