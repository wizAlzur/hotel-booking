// Утилиты для localStorage
const BookingStorage = {
    save(data) {
        localStorage.setItem('hotelBooking', JSON.stringify(data));
    },
    load() {
        return JSON.parse(localStorage.getItem('hotelBooking')) || {};
    },
    clear() {
        localStorage.removeItem('hotelBooking');
    }
};