package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello with string", func(t *testing.T) {
		got := Hello("Foo")
		want := "Hello, Foo!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello without string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})
}

// For helper functions, it's a good idea to accept a `testing.TB` so that helper functions from
// both a test or a benchmark can be called.
func assertCorrectMessage(t testing.TB, got, want string) {
	// Tells the test suite that this method is a helper.
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
