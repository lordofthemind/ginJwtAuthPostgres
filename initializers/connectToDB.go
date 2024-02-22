package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	elephantConnection := os.Getenv("ELEPHANT_CONNECTION")
	DB, err = gorm.Open(postgres.Open(elephantConnection), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
}
