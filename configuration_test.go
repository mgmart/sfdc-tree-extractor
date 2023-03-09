package main

import "testing"

func TestConfig(t *testing.T) {

	got := getConfiguration()
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
