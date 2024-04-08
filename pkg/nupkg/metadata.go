package nupkg

import "gorm.io/gorm"

type Metadata struct {
	gorm.Model
	DatabaseId               uint
	Id                       int
	Version                  string
	Title                    string
	Authors                  string
	Owners                   string
	LicenseUrl               string
	ProjectUrl               string
	IconUrl                  string
	RequireLicenseAcceptance bool
	Description              string
	ReleaseNotes             string
	Tags                     []Tag        `gorm:"foreignKey:DatabaseId"`
	Dependencies             []Dependency `gorm:"foreignKey:DatabaseId"`
}
