package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	testEnv := os.Getenv("TEST_ENV")
	if testEnv == "" {
		panic("TEST_ENV not set")
	}

}
