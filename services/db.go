package services

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func init() {
	dbinfo := "host=" + Config.PostgresHost +
		" dbname=" + Config.PostgresDatabase +
		" user=" + Config.PostgresUser +
		" password=" + Config.PostgresPassword +
		" sslmode="

	if Config.PostgresSSL {
		dbinfo += "enable"
	} else {
		dbinfo += "disable"
	}

	db, err := gorm.Open("postgres", dbinfo)

	if err != nil {
		panic("failed to connect to database")
	}

	DB = db

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}
