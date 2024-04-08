package nupkg

import "gorm.io/gorm"

type Package struct {
	gorm.Model
	DatabaseId uint     `gorm:"primaryKey"`
	Metadata   Metadata `gorm:"foreignKey:DatabaseId"`
}
