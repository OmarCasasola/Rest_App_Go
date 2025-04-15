package handlers

import (
	"GIN_GORM/db"
	"GIN_GORM/models"
	"encoding/json"
	"net/http"
)

func GetArtists(rw http.ResponseWriter) {
	var artist []models.Artist
	result := db.Database.Find(&artist)
	if result.Error != nil {
		sendError(rw, http.StatusInternalServerError, result.Error)
		return
	}
	sendData(rw, http.StatusOK, artist)
}

func SetArtists(rw http.ResponseWriter, req *http.Request) {
	var artist models.Artist
	err := json.NewDecoder(req.Body).Decode(&artist)
	if err != nil {
		sendError(rw, http.StatusBadRequest, err)
		return
	}

	result := db.Database.Save(&artist)
	if result.Error != nil {
		sendError(rw, http.StatusInternalServerError, result.Error)
		return
	}
	sendData(rw, http.StatusOK, artist)
}

func DeleteArtists(rw http.ResponseWriter, req *http.Request) {
	var artist models.Artist
	err := json.NewDecoder(req.Body).Decode(&artist)
	if err != nil {
		sendError(rw, http.StatusBadRequest, err)
		return
	}

	result := db.Database.Delete(&artist)
	if result.Error != nil {
		sendError(rw, http.StatusInternalServerError, result.Error)
		return
	}
	sendData(rw, http.StatusOK, artist)
}
