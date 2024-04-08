package nupkg

import "gorm.io/gorm"

type Dependency struct {
	gorm.Model
	DatabaseId uint
	Id         int
	Version    string
}
