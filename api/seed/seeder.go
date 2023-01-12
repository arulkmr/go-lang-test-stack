package seed

import (
	"log"

	"go-lang-test-stack/api/models"

	"github.com/jinzhu/gorm"
)

var locations = []models.Location{
	{
		LocationId: "LocId229",
		CustomerId: "hexauuidcustoomer1&9",
		Address:    "SBO Bangalore",
		Lat:        10.1,
		Long:       100.2,
	},
	{
		LocationId: "LocId334",
		CustomerId: "hexauuidcustoomer29&i",
		Address:    "Shell Chennai",
		Lat:        2.21,
		Long:       200.2,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Location{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Location{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range locations {
		err = db.Debug().Model(&models.Location{}).Create(&locations[i]).Error
		if err != nil {
			log.Fatalf("cannot seed locations table: %v", err)
		}
	}
}
