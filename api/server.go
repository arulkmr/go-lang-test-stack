package api

import (
	"fmt"
	"log"
	"os"

	"go-lang-test-stack/api/db"
	"go-lang-test-stack/api/routes"
	"go-lang-test-stack/api/seed"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	db.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	routes.InitializeRoutes(os.Getenv("SERVER_PORT"))
	seed.Load(db.DB.Db)

}
