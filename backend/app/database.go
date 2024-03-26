package app

import (
	"blogv2/database"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	UserName     string
	Password     string
	IPAddress    string
	Port         string
	DatabaseName string
	DB           *gorm.DB
}

func (config *DatabaseConfig) getConfigFile() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file, Err: %s", err)
	}
	//after config with env file or docker compose
	fmt.Println(os.Getenv("DATABASE_USER"))
	config.UserName = os.Getenv("DATABASE_USER")
	config.Password = os.Getenv("DATABASE_PASSWORD")
	config.IPAddress = os.Getenv("DATABASE_IP")
	config.Port = os.Getenv("DATABASE_PORT")
	config.DatabaseName = os.Getenv("DATABASE_NAME")
}

func (config *DatabaseConfig) Connect() {
	config.getConfigFile()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.UserName, config.Password, config.IPAddress, config.Port, config.DatabaseName)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database: " + err.Error())
	}
	config.DB = db
}

func (config *DatabaseConfig) Migration() {
	database.Migration(config.DB)
}
