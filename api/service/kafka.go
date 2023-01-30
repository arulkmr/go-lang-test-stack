package service

import (
	"encoding/json"
	"fmt"
	"go-lang-test-stack/api/db"
	"go-lang-test-stack/api/models"
)

func KafkaConsumer() {

	// Subscribe to the topic
	err := db.DB.KafkaConsumer.SubscribeTopics([]string{"location"}, nil)
	if err != nil {
		panic(err)
	}

	// Continuously poll for new messages
	for {
		msg, err := db.DB.KafkaConsumer.ReadMessage(-1)
		if err == nil {
			var myData models.Location
			err := json.Unmarshal(msg.Value, &myData)
			if err != nil {
				fmt.Printf("Error decoding JSON: %v\n", err)
				continue
			}
			fmt.Printf("Received data: Name: %s", myData.CustomerName, myData)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}

	db.DB.KafkaConsumer.Close()
}
