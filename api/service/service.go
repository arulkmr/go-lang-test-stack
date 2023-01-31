package service

import (
	"fmt"
	"go-lang-test-stack/api/db"
	"go-lang-test-stack/api/models"
	"math"

	"gorm.io/gorm"
)

func SaveLocation(l *models.Location) (*models.Location, error) {
	l.LocationId = l.GenerateId()
	err := db.DB.Db.Debug().Create(&l).Error
	if err != nil {
		return l, err
	}

	// jsonPayload, _ := json.Marshal(&l)
	// topic := "location"
	// delivery_chan := make(chan kafka.Event, 10000)
	// err = db.DB.KafkaProducer.Produce(&kafka.Message{
	// 	TopicPartition: kafka.TopicPartition{
	// 		Topic:     &topic,
	// 		Partition: kafka.PartitionAny,
	// 	},
	// 	Value: []byte(jsonPayload)},
	// 	delivery_chan,
	// )
	if err != nil {
		fmt.Println("Error producing Kafka message: " + err.Error())
	}

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
	return &location, err
}

func UpdateALocation(locId string, l *models.Location) (*models.Location, error) {

	var location = models.Location{}

	data := db.DB.Db.Debug().Model(&models.Location{}).Where("location_id = ?", locId).Take(&models.Location{}).UpdateColumns(
		map[string]interface{}{
			"Address":      l.Address,
			"CustomerId":   l.CustomerId,
			"CustomerName": l.CustomerName,
			"LocationName": l.LocationName,
			"Long":         l.Long,
			"Lat":          l.Lat,
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

// func LocationQuery(l *models.LocationQuery) (*[]models.Location, error) {
// 	var location = []models.Location{}
// 	fmt.Println("filters", l)
// 	err := db.DB.Db.Debug().Model(models.Location{}).Where("location_id = ? or customer_id=?", l.SearchbyId, l.SearchbyCustomerId).Take(&location).Error
// 	if err != nil {
// 		return &[]models.Location{}, err
// 	}
// 	return &location, err
// }

func LocationQuery(l models.LocationQuery) (*models.LocationPagination, error) {
	var locations models.LocationPagination
	var location []*models.Location

	result := db.DB.Db.Session(&gorm.Session{})

	if len(l.CustomerNames) > 0 {
		result = result.Where("customer_name IN ?", l.CustomerNames)
	}
	if len(l.LocationNames) > 0 {
		result = result.Where("location_name IN ?", l.LocationNames)
	}

	// paginate and sort
	result = result.Scopes(Paginate(&location, &locations, l, result)).Find(&location)
	locations.Locations = &location

	if result.Error != nil {
		return nil, result.Error
	}

	return &locations, nil
}

func Paginate(location *[]*models.Location, l *models.LocationPagination, qp models.LocationQuery, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	l.Limit = qp.Limit
	l.Page = qp.Page

	var totalRows int64
	db.Model(location).Count(&totalRows)
	l.TotalRows = totalRows

	totalPages := int(math.Ceil(float64(totalRows) / float64(l.GetLimit())))
	l.TotalPages = totalPages

	if qp.Sort {
		l.Sort = "customer_name, location_name"
		return func(db *gorm.DB) *gorm.DB {
			return db.Offset(l.GetOffset()).Limit(l.GetLimit()).Order("customer_name").Order("location_name")
		}
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(l.GetOffset()).Limit(l.GetLimit())
	}
}
