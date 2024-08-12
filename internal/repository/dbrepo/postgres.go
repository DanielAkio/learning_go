package dbrepo

import (
	"context"
	"errors"
	"time"

	"github.com/DanielAkio/learning_go/internal/models"
	"golang.org/x/crypto/bcrypt"
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
			AND room_id = $1
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

func (m *postgresDBRepo) GetUserByID(id int) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			password,
			access_level,
			created_at,
			updated_at
		FROM
			users
		WHERE id = $1
	`

	row := m.DB.QueryRowContext(ctx, query, id)

	var u models.User
	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.AccessLevel,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return u, err
	}

	return u, nil
}

func (m *postgresDBRepo) UpdatedUser(u models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		UPDATE users
		SET first_name=$2, last_name=$3, email=$4, access_level=$5, updated_at=$6
		WHERE id = $1
	`

	_, err := m.DB.ExecContext(ctx, query,
		u.ID,
		u.FirstName,
		u.LastName,
		u.Email,
		u.AccessLevel,
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *postgresDBRepo) Authenticate(email, testPassword string) (int, string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT id, password
		FROM users
		WHERE email = $1
	`
	row := m.DB.QueryRowContext(ctx, query, email)

	var id int
	var hashedPassword string
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		return id, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(testPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, "", errors.New("incorrect password")
	} else if err != nil {
		return 0, "", err
	}

	return id, hashedPassword, nil
}

func (m *postgresDBRepo) AllReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var reservations []models.Reservation

	query := `
		SELECT
			r.id,
			r.first_name,
			r.last_name,
			r.email,
			r.phone,
			r.start_date,
			r.end_date,
			r.room_id,
			r.created_at,
			r.updated_at,
			r.processed,
			rm.id,
			rm.name
		FROM reservations AS r
		LEFT JOIN rooms rm ON r.room_id = rm.id
		ORDER BY r.start_date ASC
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return reservations, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Processed,
			&i.Room.ID,
			&i.Room.Name,
		)

		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, i)
	}

	if err = rows.Err(); err != nil {
		return reservations, err
	}

	return reservations, nil
}

func (m *postgresDBRepo) AllNewReservations() ([]models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var reservations []models.Reservation

	query := `
		SELECT
			r.id,
			r.first_name,
			r.last_name,
			r.email,
			r.phone,
			r.start_date,
			r.end_date,
			r.room_id,
			r.created_at,
			r.updated_at,
			r.processed,
			rm.id,
			rm.name
		FROM reservations AS r
		LEFT JOIN rooms rm ON r.room_id = rm.id
		WHERE r.processed = 0
		ORDER BY r.start_date ASC
	`

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return reservations, err
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Reservation
		err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.Phone,
			&i.StartDate,
			&i.EndDate,
			&i.RoomID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Processed,
			&i.Room.ID,
			&i.Room.Name,
		)

		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, i)
	}

	if err = rows.Err(); err != nil {
		return reservations, err
	}

	return reservations, nil
}

func (m *postgresDBRepo) GetReservationByID(id int) (models.Reservation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var res models.Reservation

	query := `
		SELECT
			r.id,
			r.first_name,
			r.last_name,
			r.email,
			r.phone,
			r.start_date,
			r.end_date,
			r.room_id,
			r.created_at,
			r.updated_at,
			r.processed,
			rm.id,
			rm.name
		FROM reservations AS r
		LEFT JOIN rooms rm ON r.room_id = rm.id
		WHERE r.id = $1
	`

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&res.ID,
		&res.FirstName,
		&res.LastName,
		&res.Email,
		&res.Phone,
		&res.StartDate,
		&res.EndDate,
		&res.RoomID,
		&res.CreatedAt,
		&res.UpdatedAt,
		&res.Processed,
		&res.Room.ID,
		&res.Room.Name,
	)

	if err != nil {
		return res, err
	}

	return res, nil
}

func (m *postgresDBRepo) UpdateReservation(res models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		UPDATE reservations
		SET
			first_name=$2,
			last_name=$3,
			email=$4,
			phone=$5,
			updated_at=$6
		WHERE id = $1
	`

	_, err := m.DB.ExecContext(ctx, query,
		res.ID,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *postgresDBRepo) DeleteReservation(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		DELETE
		FROM reservations
		WHERE id = $1
	`

	_, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (m *postgresDBRepo) UpdateProcessedForReservation(id, processed int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		UPDATE reservations
		SET processed=$2
		WHERE id = $1
	`

	_, err := m.DB.ExecContext(ctx, query, id, processed)
	if err != nil {
		return err
	}

	return nil
}
