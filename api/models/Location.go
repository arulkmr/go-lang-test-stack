package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
)

type Location struct {
	LocationId string    `json:"locationid" gorm:"primaryKey"`
	CustomerId string    `gorm:"size:100;not null;unique" json:"customerid"`
	Address    string    `gorm:"size:100;not null;" json:"address"`
	Lat        float64   `gorm:"size:100;not null;" json:"lat"`
	Long       float64   `gorm:"size:100;not null;" json:"long"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c Location) GenerateId() string {
	keyString := fmt.Sprintf("%s-%s-%s", c.CustomerId, c.Address, c.Address)
	return GetMD5Hash(keyString)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GetMD5Hash(text string) string {
	// create a random string to prepend
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	// make hash
	hash := md5.Sum([]byte(string(b) + text))
	return hex.EncodeToString(hash[:])
}

func (u *Location) SaveLocation(db *gorm.DB) (*Location, error) {

	u.LocationId = u.GenerateId()
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Location{}, err
	}
	return u, nil
}

func (u *Location) FindAllLocations(db *gorm.DB) (*[]Location, error) {
	var err error
	location := []Location{}
	err = db.Debug().Model(&Location{}).Limit(100).Find(&location).Error
	if err != nil {
		return &[]Location{}, err
	}
	return &location, err
}

func (u *Location) FindLocationByID(db *gorm.DB, locId string) (*Location, error) {
	var err error
	err = db.Debug().Model(Location{}).Where("location_id = ?", locId).Take(&u).Error
	if err != nil {
		return &Location{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Location{}, errors.New("Location Not Found")
	}
	return u, err
}

func (u *Location) UpdateALocation(db *gorm.DB, locId string) (*Location, error) {

	db = db.Debug().Model(&Location{}).Where("location_id = ?", locId).Take(&Location{}).UpdateColumns(
		map[string]interface{}{
			"Address":    u.Address,
			"CustomerId": u.CustomerId,
			"Long":       u.Long,
			"Lat":        u.Lat,
		},
	)
	if db.Error != nil {
		return &Location{}, db.Error
	}
	// This is the display the updated location
	err := db.Debug().Model(&Location{}).Where("location_id = ?", locId).Take(&u).Error
	if err != nil {
		return &Location{}, err
	}
	return u, nil
}

func (u *Location) DeleteALocation(db *gorm.DB, locId string) (int64, error) {

	db = db.Debug().Model(&Location{}).Where("location_id = ?", locId).Take(&Location{}).Delete(&Location{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
