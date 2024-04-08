package internal

import (
	_ "github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gen"
	_ "gorm.io/gen"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// TODO: Implement database access
func RetrieveDatabaseConnection() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	DB, err := gorm.Open(sqlite.Open("nugo.db"), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	g.UseDB(DB)
}
