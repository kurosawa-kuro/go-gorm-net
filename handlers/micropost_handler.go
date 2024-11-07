package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go-gorm-net/database"
	"go-gorm-net/models"
)

func HandleMicroposts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var microposts []models.Micropost
		database.DB.Find(&microposts)
		json.NewEncoder(w).Encode(microposts)

	case http.MethodPost:
		var micropost models.Micropost
		json.NewDecoder(r.Body).Decode(&micropost)
		database.DB.Create(&micropost)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(micropost)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleMicropost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/microposts/")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var micropost models.Micropost
	idInt, _ := strconv.Atoi(id)
	result := database.DB.First(&micropost, idInt)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(micropost)
}
