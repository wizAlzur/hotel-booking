package models

import (
	"errors"
	"github.com/google/uuid"
)

type RoomType string

var (
	ErrRoomNotAvailable = errors.New("номер уже забронирован на выбранные даты")
)

const (
	RoomStandard  RoomType = "standard"
	RoomLux       RoomType = "lux"
	RoomPresident RoomType = "president"
)

type Booking struct {
	ID         uuid.UUID `json:"id"`
	RoomType   RoomType  `json:"room_type"`
	CheckIn    string    `json:"check_in"` // "2025-10-29"
	CheckOut   string    `json:"check_out"`
	GuestName  string    `json:"guest_name"`
	GuestEmail string    `json:"guest_email"`
	GuestPhone string    `json:"guest_phone"`
	CardNumber string    `json:"card_number"`
	CardExpiry string    `json:"card_expiry"`
	CardCVV    string    `json:"card_cvv"`
	CardHolder string    `json:"card_holder"`
}

type CreateBookingRequest struct {
	RoomType   RoomType `json:"room_type" validate:"required,oneof=standard lux president"`
	CheckIn    string   `json:"check_in" validate:"required,date"`
	CheckOut   string   `json:"check_out" validate:"required,date,gtfield=CheckIn"`
	GuestName  string   `json:"guest_name" validate:"required"`
	GuestEmail string   `json:"guest_email" validate:"required,email"`
	GuestPhone string   `json:"guest_phone" validate:"required"`
	CardNumber string   `json:"card_number" validate:"required,len=19"`
	CardExpiry string   `json:"card_expiry" validate:"required,len=5"`
	CardCVV    string   `json:"card_cvv" validate:"required,len=3"`
	CardHolder string   `json:"card_holder" validate:"required"`
}
