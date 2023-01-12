package db

import (
	"fmt"
	"go-lang-test-stack/api/models"
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql database driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Initialize(Dbdriver, DbLocation, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbLocation, DbName, DbPassword)
	db, err := gorm.Open(Dbdriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	db.Debug().AutoMigrate(&models.Location{})

	DB = Dbinstance{
		Db: db,
	}
}
