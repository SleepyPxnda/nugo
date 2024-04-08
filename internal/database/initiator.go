package database

import (
	_ "github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	_ "gorm.io/gen"
	"gorm.io/gorm"
	"server/pkg/nupkg"
	_ "server/pkg/nupkg"
)

var DB *gorm.DB

func Migrate() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	DB.AutoMigrate(&nupkg.Dependency{})
	DB.AutoMigrate(&nupkg.Metadata{})
	DB.AutoMigrate(&nupkg.Package{})
}
