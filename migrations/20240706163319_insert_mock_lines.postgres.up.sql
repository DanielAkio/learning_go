INSERT INTO rooms (name, created_at, updated_at)
VALUES ('Casa do Daniel', '2024-01-01', '2024-01-01');

INSERT INTO rooms (name, created_at, updated_at)
VALUES ('Casa da Cata', '2024-01-01', '2024-01-01');

INSERT INTO restrictions (restriction_name, created_at, updated_at)
VALUES ('Reservation', '2024-01-01', '2024-01-01');

INSERT INTO restrictions (restriction_name, created_at, updated_at)
VALUES ('Owner Block', '2024-01-01', '2024-01-01');

INSERT INTO reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at)
VALUES ('Daniel', 'Akio', 'danielakioteixeira@gmail.com', '56612347', '2024-01-01', '2024-01-10', '1', '2024-01-01', '2024-01-01');

INSERT INTO room_restrictions (start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at)
VALUES ('2024-07-01', '2024-07-31', '1', '1', '1', '2024-01-01', '2024-01-01');