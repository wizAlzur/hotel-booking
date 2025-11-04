document.addEventListener('DOMContentLoaded', () => {
    const today = new Date().toISOString().split('T')[0];
    document.getElementById('checkin').min = today;
    document.getElementById('checkout').min = today;

    document.getElementById('check-availability').addEventListener('click', () => {
        const checkin = document.getElementById('checkin').value;
        const checkout = document.getElementById('checkout').value;

        if (!checkin || !checkout || checkin >= checkout) {
            alert('Выберите корректные даты заезда и выезда');
            return;
        }

        BookingStorage.save({ checkin, checkout });
        location.href = 'rooms.html';
    });
});