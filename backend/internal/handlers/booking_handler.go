package handlers

import (
	"encoding/json"
	"hotel-prestige-backend/internal/models"
	"hotel-prestige-backend/internal/service"
	"hotel-prestige-backend/pkg/utils"
	"net/http"
)

type BookingHandler struct {
	service *service.BookingService
}

func NewBookingHandler(service *service.BookingService) *BookingHandler {
	return &BookingHandler{service: service}
}

func (h *BookingHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req models.CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if req.CheckIn >= req.CheckOut {
		utils.Error(w, http.StatusBadRequest, "Check-out must be after check-in")
		return
	}

	booking, err := h.service.Create(&req)
	if err != nil {
		if err == models.ErrRoomNotAvailable {
			utils.Error(w, http.StatusConflict, "Номер уже забронирован на выбранные даты")
		} else {
			utils.Error(w, http.StatusInternalServerError, "Ошибка сервера")
		}
		return
	}

	utils.Success(w, map[string]string{
		"booking_id": booking.ID.String(),
		"message":    "Бронь успешно создана",
	})
}
