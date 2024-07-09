package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("hello with string", func(t *testing.T) {
		got := Hello("Foo", "English")
		want := "Hello, Foo!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("hello without string", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("hello in spanish", func(t *testing.T) {
		got := Hello("Bar", "Spanish")
		want := "Hola, Bar!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("hello in french", func(t *testing.T) {
		got := Hello("Bar", "French")
		want := "Bonjour, Bar!"

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
