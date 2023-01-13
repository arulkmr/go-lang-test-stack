package service

import (
	"errors"
	"go-lang-test-stack/api/db"
	"go-lang-test-stack/api/models"

	"github.com/jinzhu/gorm"
)

type ServiceLogic interface {
	SaveLocation(l *models.Location) (*models.Location, error)
	FindAllLocations() (*models.Location, error)
	FindLocationByID(string) (*models.Location, error)
	UpdateALocation(id string, l models.Location) (*models.Location, error)
	DeleteALocation(id string) (int64, error)
}

type Database struct {
	db.Dbinstance
}

func (d *Database) SaveLocation(l *models.Location) (*models.Location, error) {
	l.LocationId = l.GenerateId()
	_ = d.Db.Debug().Create(&l).Error
	return l, nil
}

func (d *Database) FindAllLocations() (*[]models.Location, error) {

	location := []models.Location{}

	err := d.Db.Debug().Model(&models.Location{}).Limit(100).Find(&location).Error
	if err != nil {
		return &[]models.Location{}, err
	}
	return &location, err
}

func (d *Database) FindLocationByID(locId string) (*models.Location, error) {
	var location = models.Location{}
	err := d.Db.Debug().Model(models.Location{}).Where("location_id = ?", locId).Take(&location).Error
	if err != nil {
		return &models.Location{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &models.Location{}, errors.New(" models.Location Not Found")
	}
	return &location, err
}

func (d *Database) UpdateALocation(locId string, l *models.Location) (*models.Location, error) {

	var location = models.Location{}

	data := d.Db.Debug().Model(&models.Location{}).Where("location_id = ?", locId).Take(&models.Location{}).UpdateColumns(
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
	err := d.Db.Debug().Model(&models.Location{}).Where("location_id = ?", locId).Take(&location).Error
	if err != nil {
		return &models.Location{}, err
	}
	return &location, nil
}

func (d *Database) DeleteALocation(locId string) (int64, error) {

	data := d.Db.Debug().Model(&models.Location{}).Where("location_id = ?", locId).Take(&models.Location{}).Delete(&models.Location{})

	if data.Error != nil {
		return 0, data.Error
	}
	return data.RowsAffected, nil
}
