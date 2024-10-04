package models

import (
	"time"
)

type Course struct {
	ID          int       `gorm:"primary_key;column:id;type:int;not null"`
	Name        string    `gorm:"column:name;type:varchar(255);not null"`
	Description string    `gorm:"column:description;type:text"`
	Link        string    `gorm:"column:link;type:varchar(255)"`
	Rating      int       `gorm:"column:rating;type:int"`
	Duration    time.Time `gorm:"column:duration;type:time"`
	Author      string    `gorm:"column:author;type:varchar(100)"`
	Institution string    `gorm:"column:institution;type:varchar(100)"`
	IDImage     int       `gorm:"index;column:id_image;type:int"`
}

type CourseResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Rating      int       `json:"rating"`
	Duration    time.Time `json:"duration"`
	Author      string    `json:"author"`
	Institution string    `json:"institution"`
	IDImage     int       `json:"id_image"`
}