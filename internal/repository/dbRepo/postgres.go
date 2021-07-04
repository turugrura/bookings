package dbrepo

import (
	"context"
	"time"

	"github.com/turugrura/bookings/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertReservation(model models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	smtm := `insert into reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
			values ($1,$2,$3,$4,$5,$6,$7,$8,$9) returning id`

	var newId int
	err := m.DB.QueryRowContext(ctx, smtm,
		model.FirstName,
		model.LastName,
		model.Email,
		model.Phone,
		model.StartDate,
		model.EndDate,
		model.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newId)

	if err != nil {
		return 0, err
	}

	return newId, nil
}

func (m *postgresDBRepo) InsertRoomRestriction(model models.RoomRestriction) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	smtm := `insert into room_restrictions (start_date, end_date, room_id, reservation_id, created_at, updated_at, restriction_id)
			values ($1,$2,$3,$4,$5,$6,$7) 
			returning id`

	var newId int
	err := m.DB.QueryRowContext(ctx, smtm,
		model.StartDate,
		model.EndDate,
		model.RoomID,
		model.ReservationID,
		time.Now(),
		time.Now(),
		model.RestrictionID,
	).Scan(&newId)

	if err != nil {
		return 0, err
	}

	return 0, nil
}
