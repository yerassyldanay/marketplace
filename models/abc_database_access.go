package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"marketplace/utils"
	"os"
)

var db *gorm.DB

func init() {
	var nameOfFunc = "DB_CONNECT"
	var err = godotenv.Load()

	if err != nil {
		utils.LogMini(&utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		err.Error(),
			Req:          		nil,
		})
	}

	var dbUsername = os.Getenv("POSTGRES_USER")
	var dbPassword = os.Getenv("POSTGRES_PASSWORD")
	var dbName = os.Getenv("POSTGRES_DB")
	var dbHost = os.Getenv("db_host")
	var dbPort = os.Getenv("db_port")

	var dbUri = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUsername, dbName, dbPassword)
	fmt.Println("Connecting Postgres using following uri:", dbUri)

	var dbConnection, dbErr = gorm.Open("postgres", dbUri)
	if dbErr != nil {
		utils.LogMini(&utils.LogMessage{
			Success:      		false,
			FunctionName: 		nameOfFunc,
			Message:      		dbErr.Error(),
			Req:          		nil,
		})
	}

	db = dbConnection

	//db.DropTableIfExists()
	db.Debug().AutoMigrate(&CourseDocuments{})
}

func GetDB () *gorm.DB {
	return db
}

