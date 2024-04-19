package main

import "testing"

func TestRun(t *testing.T) {
	err := run()
	if err != nil {
		t.Error("failed run()")
	}
}

// Running test
// cd cmd/web
// go test -v

// Coverage
// go test -cover
// go test -coverprofile=coverage.out && go tool cover -html=coverage.out
