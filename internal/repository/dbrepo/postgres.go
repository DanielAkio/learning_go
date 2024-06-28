package dbrepo

import (
	"context"
	"time"

	"github.com/DanielAkio/learning_go/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertReservation(re models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `
		INSERT INTO reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id
	`

	_, err := m.DB.ExecContext(
		ctx,
		stmt,
		re.FirstName,
		re.LastName,
		re.Email,
		re.Phone,
		re.StartDate,
		re.EndDate,
		re.RoomID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}

	return nil
}
