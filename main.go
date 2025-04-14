package main

import (
	"GIN_GORM/db"
	"GIN_GORM/models"
)

func main() {
	db.ConnectDatabase()
	models.ArtistMigration()
	models.RoleMigration()
	models.MemberMigration()
	models.AlbumMigration()
	models.SongMigration()
}
