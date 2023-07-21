package ispanicking_test

import (
	"testing"

	ispanicking "github.com/profmagija/is-panicking"
)

func TestIsNotPanicking(t *testing.T) {
	defer func() {
		if ispanicking.IsPanicking() {
			t.Error("Should not be panicking")
		}
	}()
}

func TestIsPanicking(t *testing.T) {
	defer func() {
		if !ispanicking.IsPanicking() {
			t.Error("Should be panicking")
		}

		val := recover()
		if val != "hehe" {
			t.Error("Should have recovered panic")
		}
	}()

	panic("hehe")
}

func TestIsPanickingButRecovered(t *testing.T) {
	defer func() {
		val := recover()
		if val != "hehe" {
			t.Error("Should have recovered panic")
		}

		if ispanicking.IsPanicking() {
			t.Error("Should not be panicking")
		}
	}()

	panic("hehe")
}
