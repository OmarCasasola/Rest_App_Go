package models

import "GIN_GORM/db"

func (Song) TableName() string {
	return "songs"
}

type Song struct {
	ID       int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	AlbumID  int64  `json:"album_id" gorm:"not null;index"`
	Title    string `json:"title" gorm:"type:varchar(255);not null"`
	Duration int64  `json:"duration" gorm:"type:bigint;not null"` // Duration in seconds
	Album    Album  `json:"albums" gorm:"foreignKey:AlbumID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func SongMigration() {
	if err := db.Database.AutoMigrate(&Song{}); err != nil {
		panic(err)
	} else {
		println("Song table migrated successfully")
	}
}
