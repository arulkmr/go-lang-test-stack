package api

import (
	"fmt"
	"log"
	"os"

	"go-lang-test-stack/api/db"
	"go-lang-test-stack/api/routes"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	db.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	routes.InitializeRoutes(os.Getenv("SERVER_PORT"))
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka-server:29092",
		"group.id":          "location_group",
		"auto.offset.reset": "smallest",
	})

	if err != nil {
		fmt.Printf("Failed to create consumer :%s\n", err)
		os.Exit(1)
	}

	err = consumer.Subscribe("location", nil)

	if err != nil {
		fmt.Printf("Failed to subscribe to topic :%s\n", err)
		os.Exit(1)
	}
	run := true
	for run {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			_, err = consumer.CommitMessage(e)
			if err == nil {
				handleMessage(e)
			}

		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		}
	}

	consumer.Close()

}

func handleMessage(msg *kafka.Message) {
	fmt.Printf("Conusmed messgae : %v", string(msg.Value))
}
