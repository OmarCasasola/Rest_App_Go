package handlers

import (
	"GIN_GORM/db"
	"GIN_GORM/models"
	"net/http"
)

func GetMembers(rw http.ResponseWriter, r *http.Request) {
	var members []models.Member
	result := db.Database.Preload("Role").Preload("Artist").Find(&members)
	if result.Error != nil {
		sendError(rw, http.StatusInternalServerError, result.Error)
		return
	}
	sendData(rw, http.StatusOK, members)
}
