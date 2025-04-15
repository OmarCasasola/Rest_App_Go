package handlers

import (
	"GIN_GORM/db"
	"GIN_GORM/models"
	"encoding/json"
	"net/http"
)

func SetRoles(rw http.ResponseWriter, req *http.Request) {
	var role models.Role
	err := json.NewDecoder(req.Body).Decode(&role)
	if err != nil {
		sendError(rw, http.StatusBadRequest, err)
		return
	}

	result := db.Database.Save(&role)
	if result.Error != nil {
		sendError(rw, http.StatusInternalServerError, result.Error)
		return
	}
	sendData(rw, http.StatusOK, role)
}
