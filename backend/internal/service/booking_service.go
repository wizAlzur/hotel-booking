package service

import (
	"github.com/google/uuid"
	"hotel-prestige-backend/internal/models"
	"hotel-prestige-backend/internal/repository"
)

type BookingService struct {
	repo *repository.BookingRepo
}

func NewBookingService(repo *repository.BookingRepo) *BookingService {
	return &BookingService{repo: repo}
}

func (s *BookingService) Create(req *models.CreateBookingRequest) (*models.Booking, error) {
	// Проверка доступности
	available, err := s.repo.IsRoomAvailable(req.RoomType, req.CheckIn, req.CheckOut)
	if err != nil {
		return nil, err
	}
	if !available {
		return nil, models.ErrRoomNotAvailable
	}

	booking := &models.Booking{
		ID:         uuid.New(),
		RoomType:   req.RoomType,
		CheckIn:    req.CheckIn,
		CheckOut:   req.CheckOut,
		GuestName:  req.GuestName,
		GuestEmail: req.GuestEmail,
		GuestPhone: req.GuestPhone,
		CardNumber: req.CardNumber,
		CardExpiry: req.CardExpiry,
		CardCVV:    req.CardCVV,
		CardHolder: req.CardHolder,
	}

	if err := s.repo.Create(booking); err != nil {
		return nil, err
	}

	return booking, nil
}
