package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Jose")
	want := "Hello, Jose"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
