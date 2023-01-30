package db

import (
	"fmt"
	"go-lang-test-stack/api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Dbinstance struct {
	Db            *gorm.DB
	KafkaProducer *kafka.Producer
	KafkaConsumer *kafka.Consumer
}

var DB Dbinstance

func Initialize(Dbdriver, DbLocation, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbLocation, DbName, DbPassword)
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	db.Debug().AutoMigrate(&models.Location{})

	producer, _ := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka-server:29092",
		"client.id":         "1000",
		"acks":              "all",
	})

	consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka-server:29092",
		"group.id":          "locGroup",
		"auto.offset.reset": "earliest",
	})

	DB = Dbinstance{
		Db:            db,
		KafkaProducer: producer,
		KafkaConsumer: consumer,
	}
}
