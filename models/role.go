package models

import "GIN_GORM/db"

func (Role) TableName() string {
	return "roles"
}

type Role struct {
	ID   int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"type:varchar(255);not null;unique"`
}

type Roles []Role

func RoleMigration() {
	if err := db.Database.AutoMigrate(&Role{}); err != nil {
		panic(err)
	} else {
		println("Role table migrated successfully")
	}
}
