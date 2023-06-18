package database

import (
	"core/models"
	"fmt"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

const dsn = "root:Root@2022@tcp(127.0.0.1:3306)/DataApi?charset=utf8mb4"

func ConnectDB() {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	logrus.SetLevel(logrus.DebugLevel)

	if err != nil {
		log.Fatal("Faileeee\n", err.Error())
		os.Exit(2)
	}
	log.Println("okeeeee")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Runninggg")

	err2 := db.AutoMigrate(&models.Collection{}, &models.Folder{}, &models.Request{}, &models.Response{})
	fmt.Println(err2)
	Database = DbInstance{Db: db}
}

//GOOS=windows GOARCH=amd64 go build -o Dat.exe .
//sudo service docker start
