package rabbitmq

import (
	"compay/core/config"
	"compay/core/eslog"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Data struct {
	Key      string
	CasinoId int64
	ModelId  int64
	EventId  int64
}

var channel *amqp.Channel
var queue amqp.Queue
var connectionErrors = make(chan error)
var connectionMutex = sync.Mutex{}
var IsStopped = false

func Init() {
	connectionMutex.Lock()
	defer connectionMutex.Unlock()

	if !config.RabbitmqEnable {
		return
	}

	connection, err := amqp.Dial("amqp://" + config.RabbitmqUsername + ":" + config.RabbitmqPassword + "@" + config.RabbitmqHost + ":" + config.RabbitmqPort + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	go func() {
		<-connection.NotifyClose(make(chan *amqp.Error))
		connectionErrors <- errors.New("connection closed")
	}()

	channel, err = connection.Channel()
	failOnError(err, "Failed to open a channel")

	queue, err = channel.QueueDeclare("es-queues", true, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	args := amqp.Table{"x-delayed-type": "direct"}
	err = channel.ExchangeDeclare("es-delayer", "x-delayed-message", true, false, false, false, args)
	failOnError(err, "Failed to declare an exchange")

	err = channel.QueueBind("es-queues", "es-queues", "es-delayer", true, nil)
	failOnError(err, "Failed to bind an exchange")
}

func failOnError(err error, msg string) {
	if err != nil {
		eslog.PanicLog(fmt.Sprintf("%s: %s", msg, err))
	}
}

func Dispatch(data Data, delayInSeconds int) {
	select {
	case err := <-connectionErrors:
		if err != nil {
			eslog.LogError("RabbitMQ disconnected, attempting to reconnect")
			Init()
		}
	default:
	}

	headers := make(amqp.Table)
	exchange := ""
	if delayInSeconds > 0 {
		headers["x-delay"] = delayInSeconds * 1000
		exchange = "es-delayer"
	}

	body, _ := json.Marshal(data)
	err := channel.PublishWithContext(
		context.Background(),
		exchange,
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
			Headers:     headers,
		})
	failOnError(err, "Failed to publish a message")
}

func CreateMessageReceiver() <-chan amqp.Delivery {
	msgs, err := channel.Consume(
		queue.Name,
		"es-queues",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")
	return msgs
}

func Stop() {
	IsStopped = true
	err := channel.Cancel("es-queues", false)
	if err != nil {
		eslog.LogError("Channel close error", err.Error())
	}
}
