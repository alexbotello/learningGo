package main

import "testing"

func testloopTenTimes(t *testing.T) {
	loopTenTimes()
	ans := loopTenTimes()
	expected := 10

	if ans != expected {
		t.Errorf("got '%d' want '%d'", ans, expected)
	}
}
