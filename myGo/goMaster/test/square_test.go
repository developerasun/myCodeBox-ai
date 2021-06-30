package main

import "testing"

func TestSquare(t *testing.T) {
	rst := square(9)
	if rst != 81 {
		t.Fail()
	}
}

func TestSquare2(t *testing.T) {
	rst := square(3)

	if rst != 9 {
		t.Errorf("squre(3) should be 9 but squre(3) returns 81")
	}
}
