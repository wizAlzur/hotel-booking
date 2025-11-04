document.addEventListener('DOMContentLoaded', () => {
    document.querySelectorAll('.btn-select').forEach(btn => {
        btn.addEventListener('click', (e) => {
            document.querySelectorAll('.room-card').forEach(card => card.classList.remove('selected'));
            const card = e.target.closest('.room-card');
            card.classList.add('selected');

            const room = {
                name: card.querySelector('h3').textContent,
                price: card.querySelector('strong').textContent
            };

            const booking = BookingStorage.load();
            booking.room = room;
            BookingStorage.save(booking);

            setTimeout(() => location.href = 'guest.html', 300);
        });
    });
});