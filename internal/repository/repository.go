package repository

import "github.com/takahiromitsui/go-web-app/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) error
}