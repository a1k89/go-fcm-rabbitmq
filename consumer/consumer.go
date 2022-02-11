package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/appleboy/go-fcm"
	"github.com/streadway/amqp"
	"log"
	"notify/firebase"

	//"notify/firebase"
	//"notify/models"
	"notify/models"
	"os"
)

func StartConsumeFromMainChannel() {
	conn, err := amqp.Dial(os.Getenv("AMQP_CONN"))
	defer conn.Close()

	if err != nil {
		CreateError(err, "Can't connect to RabbitMQ")
	}

	ch, err := conn.Channel()

	if err != nil {
		CreateError(err, "Can't connect to Channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		os.Getenv("QUEUE"),
		false,
		false,
		false,
		false,
		nil)

	if err != nil {
		CreateError(err, "Can't listen queue")
	}
	msg, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		CreateError(err, "Can't consume")
	}
	client, err := fcm.NewClient(os.Getenv("FCM_API_KEY"))
	if err != nil {
		fmt.Print("FCM credentials is not correct")
		return
	}
	forever := make(chan bool)
	go func() {
		for d := range msg {
			var payload models.MessageIn
			err := json.Unmarshal(d.Body, &payload)
			if err != nil {
				fmt.Print("Can't parse message from rabbitmq")
				return
			}
			firebase.SendPush(payload, *client)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func CreateError(err error, msg string) {
	log.Fatal(err, msg)
}
