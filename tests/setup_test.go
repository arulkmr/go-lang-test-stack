package tests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"go-lang-test-stack/api/db"
	"go-lang-test-stack/api/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("./../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	Database()

	os.Exit(m.Run())

}

func Database() {

	var err error

	TestDbDriver := os.Getenv("TEST_DB_DRIVER")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"), os.Getenv("TEST_DB_USER"), os.Getenv("TEST_DB_NAME"), os.Getenv("TEST_DB_PASSWORD"))
	db.DB.Db, err = gorm.Open(TestDbDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", TestDbDriver)
	}

}

func refreshLocationTable() error {
	err := db.DB.Db.DropTableIfExists(&models.Location{}).Error
	if err != nil {
		return err
	}
	err = db.DB.Db.AutoMigrate(&models.Location{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneLocation() (models.Location, error) {

	err := refreshLocationTable()
	if err != nil {
		log.Fatal(err)
	}

	location := models.Location{
		LocationId: "loc1",
		CustomerId: "hexauuidcustoomer1&9",
		Address:    "SBO Bangalore",
		Lat:        10.1,
		Long:       100.2,
	}

	err = db.DB.Db.Model(&models.Location{}).Create(&location).Error
	if err != nil {
		return models.Location{}, err
	}
	return location, nil
}

func seedLocations() ([]models.Location, error) {

	var err error
	if err != nil {
		return nil, err
	}
	locations := []models.Location{
		{
			CustomerId: "hexauuidcustoomer1&9",
			Address:    "SBO Bangalore",
			Lat:        10.1,
			Long:       100.2,
		},
		{
			CustomerId: "hexauuidcustoomer29&i",
			Address:    "Shell Chennai",
			Lat:        2.21,
			Long:       200.2,
		},
	}

	for i, _ := range locations {
		err := db.DB.Db.Model(&models.Location{}).Create(&locations[i]).Error
		if err != nil {
			return []models.Location{}, err
		}
	}
	return locations, nil
}
