package service

import (
	"errors"
	"fmt"
	"go-lang-test-stack/api/db"
	"go-lang-test-stack/api/models"

	"github.com/jinzhu/gorm"
)

func SaveLocation(l *models.Location) (*models.Location, error) {
	l.LocationId = l.GenerateId()
	_ = db.DB.Db.Debug().Create(&l).Error
	return l, nil
}

func FindAllLocations() (*[]models.Location, error) {

	location := []models.Location{}

	err := db.DB.Db.Debug().Model(&models.Location{}).Limit(100).Find(&location).Error
	if err != nil {
		return &[]models.Location{}, err
	}
	return &location, err
}

func FindLocationByID(locId string) (*models.Location, error) {
	var location = models.Location{}
	err := db.DB.Db.Debug().Model(models.Location{}).Where("location_id = ?", locId).Take(&location).Error
	if err != nil {
		return &models.Location{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &models.Location{}, errors.New(" models.Location Not Found")
	}
	return &location, err
}

func UpdateALocation(locId string, l *models.Location) (*models.Location, error) {

	var location = models.Location{}

	data := db.DB.Db.Debug().Model(&models.Location{}).Where("location_id = ?", locId).Take(&models.Location{}).UpdateColumns(
		map[string]interface{}{
			"Address":    l.Address,
			"CustomerId": l.CustomerId,
			"Long":       l.Long,
			"Lat":        l.Lat,
		},
	)
	if data.Error != nil {
		return &models.Location{}, data.Error
	}
	// This is the display the updated location
	err := db.DB.Db.Debug().Model(&models.Location{}).Where("location_id = ?", locId).Take(&location).Error
	if err != nil {
		return &models.Location{}, err
	}
	return &location, nil
}

func DeleteALocation(locId string) (int64, error) {

	data := db.DB.Db.Debug().Model(&models.Location{}).Where("location_id = ?", locId).Take(&models.Location{}).Delete(&models.Location{})

	if data.Error != nil {
		return 0, data.Error
	}
	return data.RowsAffected, nil
}

func LocationQuery(l *models.LocationQuery) (*[]models.Location, error) {
	var location = []models.Location{}
	fmt.Println("filters", l)
	err := db.DB.Db.Debug().Model(models.Location{}).Where("location_id = ? or customer_id=?", l.SearchbyId, l.SearchbyCustomerId).Take(&location).Error
	if err != nil {
		return &[]models.Location{}, err
	}
	return &location, err
}
