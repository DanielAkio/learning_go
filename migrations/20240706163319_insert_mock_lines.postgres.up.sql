INSERT INTO rooms (name, created_at, updated_at)
VALUES ('Casa do Daniel', '2024-08-01', '2024-08-01');

INSERT INTO rooms (name, created_at, updated_at)
VALUES ('Casa da Cata', '2024-08-01', '2024-08-01');

INSERT INTO restrictions (restriction_name, created_at, updated_at)
VALUES
('Reservation', '2024-08-01', '2024-08-01'),
('Owner Block', '2024-08-01', '2024-08-01');

INSERT INTO reservations (first_name, last_name, email, phone, start_date, end_date, room_id, created_at, updated_at, processed)
VALUES
('Daniel', 'Akio', 'daniel@gmail.com', '56612347', '2024-08-01', '2024-08-02', '1', '2024-08-01', '2024-08-01', 1),
('Sônia', 'Esther', 'sonia@gmail.com', '56612347', '2024-08-03', '2024-08-04', '2', '2024-08-01', '2024-08-01', 0),
('Isabelle', 'Sandra', 'isabelle@gmail.com', '56612347', '2024-08-05', '2024-08-06', '1', '2024-08-01', '2024-08-01', 0),
('Henry', 'Julio', 'henry@gmail.com', '56612347', '2024-08-07', '2024-08-08', '2', '2024-08-01', '2024-08-01', 0),
('João', 'Filipe', 'joao@gmail.com', '56612347', '2024-08-09', '2024-08-10', '1', '2024-08-01', '2024-08-01', 1),
('Isabel', 'Louise', 'isabel@gmail.com', '56612347', '2024-08-11', '2024-08-12', '2', '2024-08-01', '2024-08-01', 1);

INSERT INTO room_restrictions (start_date, end_date, room_id, reservation_id, restriction_id, created_at, updated_at)
VALUES
('2024-08-01', '2024-08-02', '1', '1', '1', '2024-08-01', '2024-08-01'),
('2024-08-03', '2024-08-04', '2', '2', '1', '2024-08-01', '2024-08-01'),
('2024-08-05', '2024-08-06', '1', '3', '1', '2024-08-01', '2024-08-01'),
('2024-08-07', '2024-08-08', '2', '4', '1', '2024-08-01', '2024-08-01'),
('2024-08-09', '2024-08-10', '1', '5', '1', '2024-08-01', '2024-08-01'),
('2024-08-11', '2024-08-12', '2', '6', '1', '2024-08-01', '2024-08-01'),
('2024-08-30', '2024-08-31', '1', null, '2', '2024-08-01', '2024-08-01'),
('2024-08-28', '2024-08-29', '2', null, '2', '2024-08-01', '2024-08-01');

INSERT INTO users (first_name, last_name, email, password, access_level, created_at, updated_at)
VALUES ('Daniel', 'Akio', 'admin@admin.com', '$2a$12$0GYloLnPl6o9sIYCCA0BjuvKOH0roD1fBGkIQAgkylDMlSEIjTI8.', '3', '2024-08-01', '2024-08-01');
