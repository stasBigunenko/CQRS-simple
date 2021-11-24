package createQueue

import (
	"CQRS-simple/cmd/http/myConfig"
	"CQRS-simple/pkg/models"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

func connectProducer() (*amqp.Connection, error) {

	config := myConfig.SetConfig()

	connStr := "amqp://" + config.RMQLog + ":" + config.RMQPass + "@" + config.RMQPath
	conn, err := amqp.Dial(connStr)
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	return conn, nil
}

func pushCommentToQueue(topic string, message []byte) error {

	conn, err := connectProducer()
	if err != nil {
		return err
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	// with this channel open, we can then start to interact
	// with the instance and declare Queues that we can publish and
	// subscribe to
	q, err := ch.QueueDeclare(
		topic,
		false,
		false,
		false,
		false,
		nil,
	)
	// We can print out the status of our Queue here
	// this will information like the amount of messages on
	// the readServ
	fmt.Println(q)
	// Handle any errors if we were unable to create the readServ
	if err != nil {
		fmt.Println(err)
	}

	// attempt to publish a message to the readServ!
	err = ch.Publish(
		"",
		topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Published Message to Queue")

	return nil
}

// create user handler
func (cq *CreateQueue) QueueCreateWrite(c models.Cud) error {

	inBytes, _ := json.Marshal(c)
	pushCommentToQueue("cud", inBytes)

	return nil
}

func (cq *CreateQueue) QueueCreateCache(up models.Cud) error {

	inBytes, _ := json.Marshal(up)
	pushCommentToQueue("read", inBytes)

	return nil
}

type CreateQueue struct{}
