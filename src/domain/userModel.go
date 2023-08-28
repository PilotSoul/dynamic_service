package domain

import "time"

type User struct {
	ID       int       `json:"id" gorm:"primary_key" swaggerignore:"true"`
	Name     string    `json:"name"`
	Segments []Segment `gorm:"many2many:user_segments" swaggerignore:"true"`
}

type UserSegment struct {
	UserID    int `gorm:"primaryKey"`
	SegmentID int `gorm:"primaryKey"`
	CreatedAt time.Time
}
