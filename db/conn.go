package db

import (
	"../service"
	"github.com/jinzhu/gorm"
)

func Conn() *gorm.DB {
	return service.DB
}
