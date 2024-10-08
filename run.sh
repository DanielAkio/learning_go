#!/bin/bash
docker-compose -f docker-compose.yml up --detach
echo

soda reset
echo

go build -o booking.exe ./cmd/web && ./booking.exe -dbuser=admin -dbpass=admin

# Usage
# ./run.sh