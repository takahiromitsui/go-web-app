package repository

import "github.com/takahiromitsui/go-web-app/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(roomID int, startDate, endDate string) (bool, error)
}