package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/streadway/amqp"
)

func main() {

	amqpURL := os.Getenv("AMQP_SERVER_URL")

	// create connection to RABBITMQ
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

	_, err = ch.QueueDeclare(
		"QueueService1",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/send", func(c *fiber.Ctx) error {
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(c.Query("msg")),
		}

		if err := ch.Publish(
			"",
			"QueueService1",
			false,
			false,
			message,
		); err != nil {
			return err
		}
		return nil
	})

	log.Fatal(app.Listen(":3000"))

}
