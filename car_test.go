package main

import (
	"strings"
	"testing"
)

func TestCar(t *testing.T) {
	var sb strings.Builder
	car := NewCar(&sb, "", 1)

	// get the starting string
	car.writeCar()
	startingTemplate := sb.String()
	sb.Reset()

	car.Move()

	if car.pos != 1 {
		t.Errorf("car position should change, want %d, got %d", 1, car.pos)
	}

	if sb.String() == startingTemplate {
		t.Error("written string is the same as the start, it should move")
	}
	sb.Reset()

	t.Run("if out of bounds", func(t *testing.T) {
		car.pos = 9999
		car.Move()

		if car.pos != 0 {
			t.Errorf("car position should be back at 0, got %d", car.pos)
		}

		if sb.String() != startingTemplate {
			t.Error("written string is not the same as the start")
			// t.Log(cmp.Diff(sb.String(), startingTemplate))
		}
	})
}

func BenchmarkWriteCar(b *testing.B) {
	var sb strings.Builder
	car := NewCar(&sb, "", 1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		car.writeCar()
	}
}
