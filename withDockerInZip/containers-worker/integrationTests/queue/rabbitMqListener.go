package tests

import (
	"context"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	RabbitMQURL        = "amqp://guest:guest@localhost:5672/"
	rabbitMQRoutingKey = "containers.scan.initialize-scan"
	rabbitMQQueue      = "initialize-scan"
)

func ListenToContainersEngineQueue(ctx context.Context, handleFunc func(body []byte)) {
	rabbitAddress := viper.GetString("RABBITMQ_ADDRESS")
	log.Println("RABBITMQ_ADDRESS:", rabbitAddress)

	rabbitMQConn, err := amqp.Dial(rabbitAddress)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer func(rabbitMQConn *amqp.Connection) {
		err = rabbitMQConn.Close()
		if err != nil {
			log.Fatal("Failed to open a RabbitMQ connection:", err)
		}
	}(rabbitMQConn)

	rabbitMQCh, err := rabbitMQConn.Channel()
	if err != nil {
		log.Fatal("Failed to open a RabbitMQ channel:", err)
	}
	defer func(rabbitMQCh *amqp.Channel) {
		err = rabbitMQCh.Close()
		if err != nil {
			log.Fatal("Failed to close a RabbitMQ channel:", err)
		}
	}(rabbitMQCh)

	err = rabbitMQCh.ExchangeDeclare(
		"containers.topic", "topic",
		false, false,
		false, false, nil,
	)
	if err != nil {
		log.Fatal("Failed to declare RabbitMQ exchange:", err)
	}

	_, err = rabbitMQCh.QueueDeclare(
		rabbitMQQueue, // Queue name
		false,         // Durable
		false,         // Delete when unused
		false,         // Exclusive
		false,         // No-wait
		nil,           // Arguments
	)
	if err != nil {
		log.Fatal("Failed to declare RabbitMQ queue:", err)
	}

	err = rabbitMQCh.QueueBind(
		rabbitMQQueue,
		rabbitMQRoutingKey,
		"containers.topic",
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to bind queue to exchange:", err)
	}

	// Consume messages from RabbitMQ
	msgs, err := rabbitMQCh.Consume(
		rabbitMQQueue, // Queue name
		"",            // Consumer
		true,          // Auto-Ack
		false,         // Exclusive
		false,         // No-local
		false,         // No-Wait
		nil,           // Arguments
	)
	if err != nil {
		log.Fatal("Failed to register a RabbitMQ consumer:", err)
	}

	// Use a goroutine to listen for cancellation signals
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			log.Println("Closing RabbitMQ listener.")
			return
		case <-sigCh:
			log.Println("Received termination signal. Closing RabbitMQ listener.")
			err = rabbitMQCh.Close()
			if err != nil {
				return
			}
			err = rabbitMQConn.Close()
			if err != nil {
				return
			}
			return
		}
	}()

	// Process messages until cancellation or signal is received
	for {
		select {
		case <-ctx.Done():
			log.Println("Closing RabbitMQ listener.")
			return
		case msg, ok := <-msgs:
			if !ok {
				log.Println("RabbitMQ message channel closed. Closing listener.")
				return
			}
			handleFunc(msg.Body)
		}
	}
}
