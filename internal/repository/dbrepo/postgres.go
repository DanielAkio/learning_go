package dbrepo

import (
	"context"
	"time"

	"github.com/DanielAkio/learning_go/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) InsertReservation(re models.Reservation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int

	query := `
		INSERT INTO reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id
	`

	row := m.DB.QueryRowContext(
		ctx,
		query,
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
	err := row.Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

func (m *postgresDBRepo) InsertRoomRestriction(rr models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO room_restrictions (start_date, end_date, room_id, reservation_id, created_at, updated_at, restriction_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := m.DB.ExecContext(
		ctx,
		query,
		rr.StartDate,
		rr.EndDate,
		rr.RoomID,
		rr.ReservationID,
		time.Now(),
		time.Now(),
		rr.RestrictionID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *postgresDBRepo) SearchAvailabilityByDatesAndRoomID(roomID int, start, end time.Time) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT
			COUNT(1)
		FROM
			room_restrictions
		WHERE 1=1
			AND room_id = $1;
			AND $2 < end_date AND $3 > start_date
	`

	var numRows int

	row := m.DB.QueryRowContext(
		ctx,
		query,
		roomID,
		start,
		end,
	)
	err := row.Scan(&numRows)

	if err != nil {
		return false, err
	}

	if numRows == 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room

	query := `
		SELECT
			r.id,
			r.name
		FROM
			rooms AS r
		WHERE
			r.id not in (
				SELECT
					room_id
				FROM
					room_restrictions AS rr
				WHERE
					$1 < rr.end_date AND $2 > rr.start_date
			)
	`

	rows, err := m.DB.QueryContext(
		ctx,
		query,
		start,
		end,
	)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err = rows.Scan(&room.ID, &room.Name)
		if err != nil {
			return rooms, err
		}
		rooms = append(rooms, room)
	}
	if err = rows.Err(); err != nil {
		return rooms, err
	}

	return rooms, nil
}

func (m *postgresDBRepo) GetRoomByID(id int) (models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var room models.Room
	query := `
		select
			id,
			name,
			created_at,
			updated_at
		from rooms
		where id = $1
	`

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&room.ID,
		&room.Name,
		&room.CreatedAt,
		&room.UpdatedAt,
	)

	if err != nil {
		return room, err
	}

	return room, nil
}
