package dbrepo

import (
	"errors"
	"time"

	"github.com/DanielAkio/learning_go/internal/models"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

func (m *testDBRepo) InsertReservation(re models.Reservation) (int, error) {
	if re.FirstName == "Fail" {
		return 1, errors.New("some error")
	}

	return 1, nil
}

func (m *testDBRepo) InsertRoomRestriction(rr models.RoomRestriction) error {
	if rr.RoomID > 2 {
		return errors.New("some error")
	}

	return nil
}

func (m *testDBRepo) SearchAvailabilityByDatesAndRoomID(roomID int, start, end time.Time) (bool, error) {
	return false, nil
}

func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room

	if id > 2 {
		return room, errors.New("some error")
	}

	return room, nil
}

func (m *testDBRepo) GetUserByID(id int) (models.User, error) {
	var u models.User

	return u, nil
}

func (m *testDBRepo) UpdatedUser(u models.User) error {
	return nil
}

func (m *testDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	return 0, "", nil
}
