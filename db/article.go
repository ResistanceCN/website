package db

import "time"

type Article struct {
	ID           int          `json:"id"`

	Author       User         `json:"author"`
	AuthorID     int          `json:"author_id";gorm:"not null"`

	Title        string       `json:"title";gorm:"not null"`
	Content      string       `json:"name";gorm:"not null"`
	CoverImage   string       `json:"cover_image"`
	Description  string       `json:"description"`

	Published    bool         `json:"published";gorm:"not null"`
	PublishedAt  time.Time    `json:"published_at"`

	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

func init() {
	Conn().AutoMigrate(&Article{})
}
