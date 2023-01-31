package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	Initialize("postgres", "arul.kumar", "password", "5432", "localhost", "demo")

	assert.NotNil(t, DB.Db)
	assert.NotNil(t, DB.KafkaProducer)
	assert.NotNil(t, DB.KafkaConsumer)
}
