package model

import (
	"time"
)

type Source struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null;column:name" json:"name"`
	ChapterID int64     `gorm:"column:chapter_id;not null" json:"chapter_id"`
	Active    bool      `gorm:"column:active;not null;default:false" json:"active"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	Scans     []Scan    `gorm:"foreignKey:SourceID;constraint:OnDelete:CASCADE" json:"scans"`
}

type Chapter struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Nbr       int       `gorm:"column:nbr;not null" json:"nbr"`
	Title     string    `gorm:"type:varchar(20);not null;column:title" json:"title"`
	MangaID   int64     `gorm:"column:manga_id;not null" json:"manga_id"`
	Sequence  int       `gorm:"column:sequence;not null" json:"sequence"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Scan struct {
	ID        int64     `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Image_ID  string    `gorm:"type:text;not null;column:image_id" json:"image_id"`
	Server    string    `gorm:"type:text;not null;column:server" json:"server"`
	Path      string    `gorm:"type:text;not null;column:path" json:"path"`
	Sequence  int       `gorm:"column:sequence;not null" json:"sequence"`
	SourceID  int64     `gorm:"column:source_id;not null" json:"source_id"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:CURRENT_TIMESTAMP" json:"created_at"`
}
