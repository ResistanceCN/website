package db

import (
	"time"
)

const (
	FACTION_UNSPECIFIED = 0
	FACTION_RESISTANCE = 1
	FACTION_ENLIGHTENED = 2
)

type User struct {
	ID         int        `json:"id"`

	GoogleID   string     `json:"google_id";gorm:"not null;unique"`
	Email      string     `json:"email";gorm:"not null;unique"`
	Name       string     `json:"name";gorm:"not null"`
	Faction    int8       `json:"faction";gorm:"not null"`

	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func init() {
	Conn().AutoMigrate(&User{})
}
