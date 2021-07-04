package repository

import "github.com/turugrura/bookings/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(model models.RoomRestriction) (int, error)
}
