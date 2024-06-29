package repository

import (
	"time"

	"github.com/DanielAkio/learning_go/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(re models.Reservation) (int, error)
	InsertRoomRestriction(rr models.RoomRestriction) error
	SearchAvailabilityByDatesAndRoomID(roomID int, start, end time.Time) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
}
