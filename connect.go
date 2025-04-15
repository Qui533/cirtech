package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"3d-print-inventory/config"
	"3d-print-inventory/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func ConnectToDatabase() {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_NAME"))
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.Printer{}, &model.Project{}, &model.Folder{}, &model.File{}, &model.Activity{})
	fmt.Println("Database Migrated")
}

func ConnectToRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "172.17.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
