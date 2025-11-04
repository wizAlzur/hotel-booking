CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS bookings (
                                        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    room_type VARCHAR(50) NOT NULL,
    check_in DATE NOT NULL,
    check_out DATE NOT NULL,
    guest_name VARCHAR(100) NOT NULL,
    guest_email VARCHAR(100) NOT NULL,
    guest_phone VARCHAR(20) NOT NULL,
    card_number VARCHAR(19) NOT NULL,
    card_expiry VARCHAR(5) NOT NULL,
    card_cvv VARCHAR(4) NOT NULL,
    card_holder VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
    );

CREATE INDEX IF NOT EXISTS idx_bookings_room_dates
    ON bookings (room_type, check_in, check_out);