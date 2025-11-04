document.addEventListener('DOMContentLoaded', () => {
    document.getElementById('guest-form').addEventListener('submit', (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const guest = {
            name: formData.get('name'),
            email: formData.get('email'),
            phone: formData.get('phone')
        };

        const booking = BookingStorage.load();
        booking.guest = guest;
        BookingStorage.save(booking);

        location.href = 'payment.html';
    });
});