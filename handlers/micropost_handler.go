package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go-gorm-net/models"
	"go-gorm-net/services"
)

type MicropostHandler struct {
	service *services.MicropostService
}

func NewMicropostHandler() *MicropostHandler {
	return &MicropostHandler{
		service: services.NewMicropostService(),
	}
}

func (h *MicropostHandler) HandleMicroposts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		microposts, err := h.service.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(microposts)

	case http.MethodPost:
		var micropost models.Micropost
		if err := json.NewDecoder(r.Body).Decode(&micropost); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.service.Create(&micropost); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(micropost)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *MicropostHandler) HandleMicropost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/microposts/")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	micropost, err := h.service.GetByID(idInt)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(micropost)
}
