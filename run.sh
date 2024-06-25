#!/bin/bash
docker-compose -f docker-compose.yml up --detach

echo

go build -o booking.exe ./cmd/web && ./booking.exe

# Using git bash
# ./run.sh