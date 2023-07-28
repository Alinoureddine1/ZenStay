package repository

import "github.com/Alinoureddine1/ZenStay/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertReservation(res models.Reservation) error
}
