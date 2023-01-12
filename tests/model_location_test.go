package tests

import (
	"log"
	"testing"

	"go-lang-test-stack/api/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/go-playground/assert.v1"
)

func TestSaveLocation(t *testing.T) {

	err := refreshLocationTable()
	if err != nil {
		log.Fatal(err)
	}

	newLocation := models.Location{
		LocationId: "loc1",
		CustomerId: "hexauuidcustoomer1&9",
		Address:    "SBO Bangalore",
		Lat:        10.1,
		Long:       100.2,
	}
	savedLocation, err := newLocation.SaveLocation(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the locations: %v\n", err)
		return
	}
	assert.Equal(t, newLocation.LocationId, savedLocation.LocationId)
	assert.Equal(t, newLocation.CustomerId, savedLocation.CustomerId)
	assert.Equal(t, newLocation.Address, savedLocation.Address)
}

func TestFindAllLocations(t *testing.T) {

	err := refreshLocationTable()
	if err != nil {
		log.Fatal(err)
	}

	_, err = seedLocations()
	if err != nil {
		log.Fatal(err)
	}

	locations, err := locationInstance.FindAllLocations(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the locations: %v\n", err)
		return
	}
	assert.Equal(t, len(*locations), 2)
}

func TestFindLocationByID(t *testing.T) {

	err := refreshLocationTable()
	if err != nil {
		log.Fatal(err)
	}

	location, err := seedOneLocation()
	if err != nil {
		log.Fatalf("cannot seed locations table: %v", err)
	}
	foundLocation, err := locationInstance.FindLocationByID(server.DB, location.LocationId)
	if err != nil {
		t.Errorf("this is the error getting one location: %v\n", err)
		return
	}
	assert.Equal(t, foundLocation.LocationId, location.LocationId)
	assert.Equal(t, foundLocation.CustomerId, location.CustomerId)
	assert.Equal(t, foundLocation.Address, location.Address)
}

func TestUpdateALocation(t *testing.T) {

	err := refreshLocationTable()
	if err != nil {
		log.Fatal(err)
	}

	location, err := seedOneLocation()
	if err != nil {
		log.Fatalf("Cannot seed location: %v\n", err)
	}

	locationUpdate := models.Location{
		LocationId: "loc1",
		CustomerId: "hexauuidcustoomer1&9",
		Address:    "SBO Bangalore",
		Lat:        10.1,
		Long:       100.2,
	}
	updatedLocation, err := locationUpdate.UpdateALocation(server.DB, location.LocationId)
	if err != nil {
		t.Errorf("this is the error updating the location: %v\n", err)
		return
	}
	assert.Equal(t, updatedLocation.LocationId, locationUpdate.LocationId)
	assert.Equal(t, updatedLocation.CustomerId, locationUpdate.CustomerId)
	assert.Equal(t, updatedLocation.Address, locationUpdate.Address)
}

func TestDeleteALocation(t *testing.T) {

	err := refreshLocationTable()
	if err != nil {
		log.Fatal(err)
	}

	location, err := seedOneLocation()

	if err != nil {
		log.Fatalf("Cannot seed location: %v\n", err)
	}

	isDeleted, err := locationInstance.DeleteALocation(server.DB, location.LocationId)
	if err != nil {
		t.Errorf("this is the error updating the location: %v\n", err)
		return
	}
	assert.Equal(t, int(isDeleted), 1)

}
