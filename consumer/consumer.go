package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"notify/firebase"
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

	forever := make(chan bool)
	go parseAndGo(msg)
	log.Print("Start listening from rabbitmq")
	<-forever
}

func CreateError(err error, msg string) {
	log.Fatal(err, msg)
}

func parseAndGo(msg <-chan amqp.Delivery) {
	for d := range msg {
		var payload models.MessageIn
		err := json.Unmarshal(d.Body, &payload)
		if err != nil {
			fmt.Print("Can't parse message from rabbitmq")
			return
		}
		//log.Printf("response:\n", payload.FcmTokens)
		firebase.SendPush(payload)
	}
}
