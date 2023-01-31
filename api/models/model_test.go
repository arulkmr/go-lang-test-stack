package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLocationModel(t *testing.T) {
	location := Location{
		LocationId:   "1",
		CustomerId:   "100",
		CustomerName: "Test Customer",
		LocationName: "Test Location",
		Address:      "Test Address",
		Lat:          37.7749,
		Long:         -122.4194,
		Connectors: []Connector{
			{
				ConnectorName: "Test Connector",
				ConnectorType: "Type 1",
			},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	assert.Equal(t, "1", location.LocationId, "Location ID should be equal to 1")
	assert.Equal(t, "100", location.CustomerId, "Customer ID should be equal to 100")
	assert.Equal(t, "Test Customer", location.CustomerName, "Customer Name should be equal to 'Test Customer'")
	assert.Equal(t, "Test Location", location.LocationName, "Location Name should be equal to 'Test Location'")
	assert.Equal(t, "Test Address", location.Address, "Address should be equal to 'Test Address'")
	assert.Equal(t, float64(37.7749), location.Lat, "Latitude should be equal to 37.7749")
	assert.Equal(t, float64(-122.4194), location.Long, "Longitude should be equal to -122.4194")
	assert.Equal(t, "Test Connector", location.Connectors[0].ConnectorName, "Connector Name should be equal to 'Test Connector'")
	assert.Equal(t, "Type 1", location.Connectors[0].ConnectorType, "Connector Type should be equal to 'Type 1'")
}

func TestConnectorModel(t *testing.T) {
	connector := Connector{
		ConnectorName: "Test Connector",
		ConnectorType: "Type 1",
	}

	assert.Equal(t, "Test Connector", connector.ConnectorName, "Connector Name should be equal to 'Test Connector'")
	assert.Equal(t, "Type 1", connector.ConnectorType, "Connector Type should be equal to 'Type 1'")
}
