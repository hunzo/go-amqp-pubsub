package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	amqpURL := os.Getenv("AMQP_SERVER_URL")

	con, err := amqp.Dial(amqpURL)
	if err != nil {
		panic(err)
	}
	defer con.Close()

	ch, err := con.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	messages, err := ch.Consume(
		"QueueService1",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	log.Println("Connect AMQP successfully")
	log.Println("Waiting for message..")

	forever := make(chan bool)

	go func() {
		for message := range messages {
			Sendmail()
			log.Printf(" > Recieved message : %s", message.Body)
		}
	}()

	<-forever

}

func Sendmail() {
	fmt.Printf("Sendmail Function")
}
