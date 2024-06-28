package repository

import "github.com/DanielAkio/learning_go/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(re models.Reservation) error
}
