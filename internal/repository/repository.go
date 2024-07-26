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
	GetRoomByID(id int) (models.Room, error)
	GetUserByID(id int) (models.User, error)
	UpdatedUser(u models.User) error
	Authenticate(email, testPassword string) (int, string, error)
}
