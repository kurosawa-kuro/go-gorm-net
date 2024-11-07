package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go-gorm-net/internal/models"
	"go-gorm-net/internal/services"
	"go-gorm-net/pkg/logger"
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
			logger.ErrorLogger.Printf("Failed to get microposts: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(microposts)

	case http.MethodPost:
		var micropost models.Micropost
		if err := json.NewDecoder(r.Body).Decode(&micropost); err != nil {
			logger.ErrorLogger.Printf("Failed to decode request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.service.Create(&micropost); err != nil {
			logger.ErrorLogger.Printf("Failed to create micropost: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(micropost)

	default:
		logger.ErrorLogger.Printf("Method not allowed: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *MicropostHandler) HandleMicropost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logger.ErrorLogger.Printf("Method not allowed: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/microposts/")
	if id == "" {
		logger.ErrorLogger.Print("Missing ID in request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		logger.ErrorLogger.Printf("Invalid ID format: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	micropost, err := h.service.GetByID(idInt)
	if err != nil {
		logger.ErrorLogger.Printf("Failed to get micropost with ID %d: %v", idInt, err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(micropost)
}
