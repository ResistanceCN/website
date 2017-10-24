package db

import "time"

const (
	STATUS_DRAFT = 0
	STATUS_IN_REVIEW = 1
	STATUS_PUBLISHED = 2
)

type Article struct {
	ID           int          `json:"id"`

	User         User         `json:"author"`
	UserID       int          `json:"author_id";gorm:"not null"`

	Title        string       `json:"title";gorm:"not null"`
	Content      string       `json:"name";gorm:"not null"`
	CoverImage   string       `json:"cover_image"`
	Description  string       `json:"description"`

	Status       int8         `json:"status";gorm:"not null"`
	PublishedAt  time.Time    `json:"published_at"`

	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

func init() {
	Conn().AutoMigrate(&Article{})
}
