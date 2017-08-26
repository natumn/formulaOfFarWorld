package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	answer := calc("15*5")
	if answer != 75 {
		t.Fatal("miss calculator.")
	}

	t.Log("clear test: ", answer)
}
