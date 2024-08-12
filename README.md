# Bookings and Reservations
This is the repository for my bookings and reservations project.

## Prerequisites
- Golang
 - buffalo soda
- Docker

## Migrations (soda)
- ```soda generate fizz <migration_name>```
- ```soda migrate```
- ```soda migrate down -s <number_of_migrations>```
- ```soda reset``` runs migrate down and up

## Go commands
- ```go get <lib_path>``` install libs
- ```go mod tidy``` remove unused libs
- ```go test -v ./...```
- ```go test -coverprofile=coverage.out && go tool cover -html=coverage.out``` coverage 1 file
- ```go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out``` coverage all files

## Running Local
- Git Bash
  - `./run.sh`
- Powershell
  - `./run.bat`
