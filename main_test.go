package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	answer := calculator("15*5")
	if answer == nil {
		t.Fatal("not calculator.")
	}
}
