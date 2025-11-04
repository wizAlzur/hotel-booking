package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"hotel-backend/config"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

type BookingRequest struct {
	RoomType   string `json:"room_type"`
	CheckIn    string `json:"check_in"`
	CheckOut   string `json:"check_out"`
	GuestName  string `json:"guest_name"`
	GuestEmail string `json:"guest_email"`
	GuestPhone string `json:"guest_phone"`
	CardNumber string `json:"card_number"`
	CardExpiry string `json:"card_expiry"`
	CardCVV    string `json:"card_cvv"`
	CardHolder string `json:"card_holder"`
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	db, err = connectDB(cfg)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer db.Close()

	if err := applyMigrations(db); err != nil {
		log.Printf("Migration failed: %v", err)
	}

	http.HandleFunc("/api/booking", bookingHandler)

	addr := ":" + cfg.ServerPort
	log.Printf("Сервер запущен на http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func connectDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	var conn *sql.DB
	var err error

	for attempt := 1; attempt <= 15; attempt++ {
		conn, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Printf("Open error (attempt %d): %v", attempt, err)
			time.Sleep(2 * time.Second)
			continue
		}

		if err = conn.Ping(); err == nil {
			log.Println("DB connected successfully")
			return conn, nil
		}

		log.Printf("Ping failed (attempt %d): %v", attempt, err)
		conn.Close()
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("failed to connect to DB after 15 attempts")
}

func applyMigrations(db *sql.DB) error {
	migrationPath := "/migrations/001_init.sql"
	data, err := os.ReadFile(migrationPath)
	if err != nil {
		return fmt.Errorf("read migration %s: %w", migrationPath, err)
	}
	if _, err = db.Exec(string(data)); err != nil {
		return fmt.Errorf("exec migration: %w", err)
	}
	log.Println("Migration applied: 001_init.sql")
	return nil
}

func bookingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"success":false,"message":"Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"success":false,"message":"Invalid JSON"}`, http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO bookings (
			room_type, check_in, check_out, guest_name, guest_email,
			guest_phone, card_number, card_expiry, card_cvv, card_holder
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id`

	var bookingID string
	err := db.QueryRow(query,
		req.RoomType, req.CheckIn, req.CheckOut, req.GuestName, req.GuestEmail,
		req.GuestPhone, req.CardNumber, req.CardExpiry, req.CardCVV, req.CardHolder,
	).Scan(&bookingID)

	if err != nil {
		log.Printf("DB insert error: %v", err)
		http.Error(w, `{"success":false,"message":"Database error"}`, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"data": map[string]string{
			"booking_id": bookingID,
			"message":    "Бронь успешно создана",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
