package model

import (
	"time"

	"github.com/lib/pq"
)

type Manga struct {
	ID                int64          `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Title             string         `gorm:"type:varchar(300);not null;column:title" json:"title"`
	AlternativeTitles pq.StringArray `gorm:"type:text[];column:alternative_titles" json:"alternative_titles"`
	CoverID           int64          `gorm:"column:cover_id" json:"-"`
	Slug              string         `gorm:"type:varchar(300);unique;not null;column:slug" json:"slug"`
	Synopsis          string         `gorm:"type:text;not null;column:synopsis" json:"synopsis"`
	LastChapter       *int           `gorm:"column:last_chapter" json:"last_chapter"`
	Author            *string        `gorm:"column:author" json:"author"`
	Themes            pq.StringArray `gorm:"type:text[];column:themes" json:"themes"`
	Year              *int           `gorm:"column:year;check:year >= 1900" json:"year"`
	Rating            *float64       `gorm:"column:rating;check:rating >= 0 AND rating <= 10" json:"rating"`
	CreatedAt         time.Time      `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP" json:"-"`
	Chapters          []Chapter      `gorm:"foreignKey:MangaID;constraint:OnDelete:CASCADE" json:"chapters"`
	Image             string         `gorm:"type:text;not null;column:image" json:"image"`
}

type Cover struct {
	Path      string    `gorm:"type:text;not null;colmun:path" json:"-"`
	ImageId   string    `gorm:"type:text;not null;column:image_id;unique" json:"-"`
	Server    string    `gorm:"type:text;not null;column:server" json:"-"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP" json:"-"`
}
