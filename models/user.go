package models

import (
	"time"

	"../services"
)

const (
	FACTION_UNSPECIFIED = 0
	FACTION_RESISTANCE = 1
	FACTION_ENLIGHTENED = 2
)

type User struct {
	ID         uint       `json:"id";gorm:"primary_key"`

	GoogleID   string     `json:"google_id";gorm:"not null;unique"`
	Email      string     `json:"email";gorm:"not null;unique"`
	Name       string     `json:"name";gorm:"not null"`
	Faction    int8       `json:"faction";gorm:"not null"`

	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func init() {
	services.DB.AutoMigrate(&User{})
}
