package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"bitly_backend/app/models"
	"bitly_backend/app/shared/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySQL struct {
	DB *gorm.DB
}

func SetupMySQL() *MySQL {
	fmt.Println("Setup... MySQL")

	url := config.DB.MySQL.Datasourcename

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Error,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	fmt.Println(url)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Product{}, )
	
	sql, err := db.DB()
	if err != nil {
		panic(err)
	}

	sql.SetConnMaxLifetime(5 * time.Minute)
	sql.SetConnMaxIdleTime(5 * time.Minute)
	sql.SetMaxIdleConns(40)
	sql.SetMaxOpenConns(40)

	return &MySQL{
		db,
	}
}
