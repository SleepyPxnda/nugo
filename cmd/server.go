package main

//TODO: SQLite starten und dort Infos halten -> einfacherer Dateizugriff

import (
	"server/internal/database"
	"server/internal/http"
)

func main() {
	database.Migrate()
	http.StartServer()
}
