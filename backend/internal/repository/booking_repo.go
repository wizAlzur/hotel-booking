package repository

import (
	"database/sql"
	"hotel-prestige-backend/internal/models"
)

type BookingRepo struct {
	db *sql.DB
}

func NewBookingRepo(db *sql.DB) *BookingRepo {
	return &BookingRepo{db: db}
}

// Проверка: есть ли пересечение дат для этого типа номера
func (r *BookingRepo) IsRoomAvailable(roomType models.RoomType, checkIn, checkOut string) (bool, error) {
	var count int
	query := `
        SELECT COUNT(*) FROM bookings
        WHERE room_type = $1
        AND (
            (check_in <= $2 AND check_out > $2) OR
            (check_in < $3 AND check_out >= $3) OR
            (check_in >= $2 AND check_out <= $3)
        )
    `
	err := r.db.QueryRow(query, roomType, checkIn, checkOut).Scan(&count)
	if err != nil {
		return false, err
	}
	return count == 0, nil
}

// Создание брони
func (r *BookingRepo) Create(booking *models.Booking) error {
	query := `
        INSERT INTO bookings (
            id, room_type, check_in, check_out,
            guest_name, guest_email, guest_phone,
            card_number, card_expiry, card_cvv, card_holder
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
    `
	_, err := r.db.Exec(query,
		booking.ID, booking.RoomType, booking.CheckIn, booking.CheckOut,
		booking.GuestName, booking.GuestEmail, booking.GuestPhone,
		booking.CardNumber, booking.CardExpiry, booking.CardCVV, booking.CardHolder,
	)
	return err
}
