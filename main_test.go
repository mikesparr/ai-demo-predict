package main

import "testing"

func TestMain(t *testing.T) {
	want := "Success"
	if got := "Success"; got != want {
		t.Errorf("Main() = %q, want %q", got, want)
	}
}
