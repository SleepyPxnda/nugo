package main

//TODO: SQLite starten und dort Infos halten -> einfacherer Dateizugriff

import "server/internal/database"

func main() {
	database.Migrate()
}
