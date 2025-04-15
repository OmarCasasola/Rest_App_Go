package models

import (
	"GIN_GORM/db"
)

func (Member) TableName() string {
	return "members"
}

type Member struct {
	ID       int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	ArtistID int64  `json:"artist_id" gorm:"not null;index"`
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	RoleID   int64  `json:"role_id" gorm:"not null;index"`
	Artist   Artist `json:"artists" gorm:"foreignKey:ArtistID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Role     Role   `json:"role" gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Members []Member

func MemberMigration() {
	if err := db.Database.AutoMigrate(&Member{}); err != nil {
		panic(err)
	} else {
		println("Member table migrated successfully")
	}
}
