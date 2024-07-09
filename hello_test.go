package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Foo")
	want := "Hello, Foo!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
