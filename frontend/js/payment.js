document.addEventListener('DOMContentLoaded', () => {
    const booking = BookingStorage.load();
    document.getElementById('summary').innerHTML = `
        <p><strong>Номер:</strong> ${booking.room?.name || '—'}</p>
        <p><strong>Даты:</strong> ${booking.checkin} — ${booking.checkout}</p>
        <p><strong>Гость:</strong> ${booking.guest?.name || '—'}</p>
        <p><strong>Стоимость:</strong> ${booking.room?.price || '—'}</p>
    `;

    document.getElementById('card-number').addEventListener('input', (e) => {
        let v = e.target.value.replace(/\D/g, '').slice(0, 16);
        v = v.match(/.{1,4}/g)?.join(' ') || v;
        e.target.value = v;
    });

    document.getElementById('payment-form').addEventListener('submit', async (e) => {
        e.preventDefault();

        const formData = new FormData(e.target);

        const fullBooking = {
            room_type: booking.room?.name === "Стандарт" ? "standard" :
                booking.room?.name === "Люкс" ? "lux" : "president",
            check_in: booking.checkin,
            check_out: booking.checkout,
            guest_name: booking.guest?.name || '',
            guest_email: booking.guest?.email || '',
            guest_phone: booking.guest?.phone || '',
            card_number: formData.get('card'),           // 1234 5678 9012 3456
            card_expiry: formData.get('expiry'),         // 12/27
            card_cvv: formData.get('cvv'),               // 123
            card_holder: formData.get('holder')          // IVAN IVANOV
        };

        console.log('Отправка на бэкенд:', fullBooking);

        try {
            const res = await fetch('/api/booking', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(fullBooking)
            });

            const data = await res.json();

            if (data.success) {
                alert(`Бронь успешно создана!\nID: ${data.data.booking_id}`);
                BookingStorage.clear();
                location.href = 'index.html';
            } else {
                alert(`Ошибка: ${data.message}`);
            }
        } catch (err) {
            console.error('Ошибка сети:', err);
            alert('Не удалось подключиться к серверу. Убедитесь, что бэкенд запущен.');
        }
    });
});