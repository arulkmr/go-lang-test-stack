package models

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Location struct {
	LocationId   string  `json:"locationid" gorm:"primaryKey"`
	CustomerId   string  `gorm:"size:100;not null;" json:"customerid,omitempty"`
	CustomerName string  `gorm:"size:100;not null;" json:"customername,omitempty"`
	LocationName string  `gorm:"size:100;not null;" json:"locationname,omitempty"`
	Address      string  `gorm:"size:100;not null;" json:"address,omitempty"`
	Lat          float64 `gorm:"size:100;not null;" json:"lat,omitempty"`
	Long         float64 `gorm:"size:100;not null;" json:"long,omitempty"`
	//Connectors   []Connector
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Connector struct {
	gorm.Model
	ConnectorName string `gorm:"size:100;not null;" json:"connectorname,omitempty"`
	ConnectorType string `gorm:"size:100;not null;" json:"connectortype,omitempty"`
}

type LocationQuery struct {
	gorm.Model
	CustomerNames []string `json:"customer_names,omitempty"`
	LocationNames []string `json:"location_names,omitempty"`
	Sort          bool     `json:"sort,omitempty"`
	Limit         int      `json:"limit,omitempty"`
	Page          int      `json:"page,omitempty"`
}

type LocationPagination struct {
	gorm.Model
	Limit      int          `json:"limit,omitempty"`
	Page       int          `json:"page,omitempty"`
	Sort       string       `json:"sort,omitempty"`
	TotalRows  int64        `json:"total_rows"`
	TotalPages int          `json:"total_pages"`
	Locations  *[]*Location `json:"connectors"`
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

func (loc *LocationPagination) GetOffset() int {
	return (loc.GetPage() - 1) * loc.GetLimit()
}

func (loc *LocationPagination) GetLimit() int {
	if loc.Limit == 0 {
		loc.Limit = 30
	}
	return loc.Limit
}

func (loc *LocationPagination) GetPage() int {
	if loc.Page == 0 {
		loc.Page = 1
	}
	return loc.Page
}
