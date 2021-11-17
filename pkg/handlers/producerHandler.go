package handlers

//import (
//	"CQRS-simple/cmd/http/myConfig"
//	"CQRS-simple/pkg/models"
//	"encoding/json"
//	"fmt"
//	"github.com/streadway/amqp"
//)
//
//func connectProducer() (*amqp.Connection, error) {
//
//	config := myConfig.SetConfig()
//
//	connStr := "amqp://" + config.RMQLog + ":" + config.RMQPass + "@" + config.RMQPath
//	conn, err := amqp.Dial(connStr)
//	if err != nil {
//		fmt.Println("Failed Initializing Broker Connection")
//		panic(err)
//	}
//
//	return conn, nil
//}
//
//func pushCommentToQueue(topic string, message []byte) error {
//
//	conn, err := connectProducer()
//	if err != nil {
//		return err
//	}
//
//	defer conn.Close()
//
//	ch, err := conn.Channel()
//	if err != nil {
//		fmt.Println(err)
//	}
//	defer ch.Close()
//
//	// with this channel open, we can then start to interact
//	// with the instance and declare Queues that we can publish and
//	// subscribe to
//	q, err := ch.QueueDeclare(
//		topic,
//		false,
//		false,
//		false,
//		false,
//		nil,
//	)
//	// We can print out the status of our Queue here
//	// this will information like the amount of messages on
//	// the queue
//	fmt.Println(q)
//	// Handle any errors if we were unable to create the queue
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	// attempt to publish a message to the queue!
//	err = ch.Publish(
//		"",
//		topic,
//		false,
//		false,
//		amqp.Publishing{
//			ContentType: "text/plain",
//			Body:        message,
//		},
//	)
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println("Successfully Published Message to Queue")
//
//	return nil
//}
//
//// create user handler
//func queueCreate(c models.Cud) error {
//
//	inBytes, _ := json.Marshal(c)
//	pushCommentToQueue("cud", inBytes)
//
//	return nil
//}
