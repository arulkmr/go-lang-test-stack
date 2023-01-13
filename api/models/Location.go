package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
	// "go-lang-test-stack/api/db"
)

type Location struct {
	LocationId string    `json:"locationid" gorm:"primaryKey"`
	CustomerId string    `gorm:"size:100;not null;" json:"customerid"`
	Address    string    `gorm:"size:100;not null;" json:"address"`
	Lat        float64   `gorm:"size:100;not null;" json:"lat"`
	Long       float64   `gorm:"size:100;not null;" json:"long"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// type Location struct {
// 	LocationId string    `json:"locationid" gorm:"primaryKey"`
// 	CustomerId string    `gorm:"size:100;not null;" json:"customerid,omitempty"`
// 	Address    string    `gorm:"size:100;not null;" json:"address,omitempty"`
// 	Lat        float64   `gorm:"size:100;not null;" json:"lat,omitempty"`
// 	Long       float64   `gorm:"size:100;not null;" json:"long,omitempty"`
// 	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
// 	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
// }

type LocationQuery struct {
	SearchbyId         string `json:"id,omitempty"`
	SearchbyCustomerId string `json:"customer_id,omitempty"`
	Limit              int    `json:"limit,omitempty"`
	Page               int    `json:"page,omitempty"`
	Sort               string `json:"sort,omitempty"`
	TotalPages         int    `json:"total_pages"`
}

// type Connector struct {
// 	ConnectorName string `gorm:"size:100;not null;" json:"connectorname,omitempty"`
// 	ConnectorType string `gorm:"size:100;not null;" json:"connectortype,omitempty"`
// }

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
