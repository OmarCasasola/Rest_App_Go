package models

import "GIN_GORM/db"

func (Album) TableName() string {
	return "albums"
}

type Album struct {
	ID          int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	ArtistID    int64  `json:"artist_id" gorm:"not null;index"`
	Title       string `json:"title" gorm:"type:varchar(255);not null"`
	ReleaseDate string `json:"release_date" gorm:"type:date;not null"`
	Artist      Artist `json:"artists" gorm:"foreignKey:ArtistID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func AlbumMigration() {
	if err := db.Database.AutoMigrate(&Album{}); err != nil {
		panic(err)
	} else {
		println("Album table migrated successfully")
	}
}
