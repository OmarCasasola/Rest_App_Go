package models

import "GIN_GORM/db"

func (Artist) TableName() string {
	return "artists"
}

type Artist struct {
	ID          int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string `json:"name" gorm:"type:varchar(255);not null;unique"`
	Genre       string `json:"genre" gorm:"type:varchar(255);not null"`
	Description string `json:"desc" gorm:"type:varchar(255);not null"`
}

func ArtistMigration() {
	if err := db.Database.AutoMigrate(&Artist{}); err != nil {
		panic(err)
	} else {
		println("Artist table migrated successfully")
	}
}
