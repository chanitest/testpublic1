package messages

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Checkmarx-Containers/containers-worker/internal/ContainersEngine"
	"github.com/checkmarxDev/lumologger/extendedlogger"
	amqp "github.com/rabbitmq/amqp091-go"
)

//go:generate moq -out rabbitMqProvider_test.go . RabbitMqInt
type RabbitMqInt interface {
	SendNewScanRequest(ctx context.Context, newScanRequest *ContainersEngine.NewScanRequest) error
}

type RabbitMq struct {
	logger       extendedlogger.ExtendedLogger
	rabbitUrl    string
	exchangeName string
	routingKey   string
}

func NewRabbitMq(logger *extendedlogger.ExtendedLogger, rabbitUrl, exchangeName, routingKey string) (*RabbitMq, error) {
	a := &RabbitMq{
		logger:       *logger,
		rabbitUrl:    rabbitUrl,
		exchangeName: exchangeName,
		routingKey:   routingKey,
	}
	return a, nil
}

func (rabbit *RabbitMq) SendNewScanRequest(ctx context.Context, newScanRequest *ContainersEngine.NewScanRequest) error {
	rabbit.logger.Info("rabbit url= %s", rabbit.rabbitUrl)

	conn, err := amqp.Dial(rabbit.rabbitUrl)
	if err != nil {
		rabbit.logger.Error("Failed to connect to RabbitMQ", err)
		return err
	}
	defer func(conn *amqp.Connection) {
		conErr := conn.Close()
		if conErr != nil {
			rabbit.logger.Error("Failed to close connection", err)
		}
	}(conn)

	ch, err := conn.Channel()
	if err != nil {
		rabbit.logger.Error("Failed to open a channel", err)
		return err
	}
	defer func(ch *amqp.Channel) {
		chErr := ch.Close()
		if chErr != nil {
			rabbit.logger.Error("Failed to close channel", err)
		}
	}(ch)

	err = ch.ExchangeDeclare(
		rabbit.exchangeName, "topic",
		false, false,
		false, false, nil,
	)
	if err != nil {
		rabbit.logger.Error("Failed to declare exchange", err)
		return err
	}

	body, err := json.Marshal(newScanRequest)
	if err != nil {
		rabbit.logger.Error("Failed to marshall new scan request", err)
		return err
	}

	err = ch.PublishWithContext(ctx,
		rabbit.exchangeName, rabbit.routingKey,
		false, false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		rabbit.logger.Error("Failed to publish message", err)
		return err
	}

	rabbit.logger.Info(fmt.Sprintf("Successfully Sent New Scan Request to %s queue", rabbit.routingKey))
	return nil
}
