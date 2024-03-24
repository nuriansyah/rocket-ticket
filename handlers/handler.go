package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/nuriansyah/rocket-ticket/model/dto"
	"github.com/nuriansyah/rocket-ticket/services"
)

type Handler struct {
	venueService *services.VenueService
}

func NewHandler(venueService *services.VenueService) *Handler {
	return &Handler{
		venueService: venueService,
	}
}

func (h *Handler) CreateVenue(w http.ResponseWriter, r *http.Request) {
	var req dto.VenueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" || req.Address == "" || req.Capacity <= 0 {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}
	if err := h.venueService.CreateVenue(req.Name, req.Address, req.Capacity); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
