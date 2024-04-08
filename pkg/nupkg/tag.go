package nupkg

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	DatabaseId uint `gorm:"primaryKey"`
	Content    string
}
