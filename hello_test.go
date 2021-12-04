package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func (t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Brandon", "")
		want := "Hello, Brandon"

		assertCorrectMessage(t, got, want)
	}) 

	t.Run("say 'Hello, world' when an empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
		
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Louise", "French")
		want := "Bonjour, Louise"
		assertCorrectMessage(t, got, want)
	})
}

func TestPrefix(t *testing.T) {
	assertCorrectPrefix := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("French prefix", func(t *testing.T) {
		got := greetingPrefix("French")
		want := "Bonjour, "
		assertCorrectPrefix(t, got, want)
	})

	t.Run("Spanish prefix", func(t *testing.T) {
		got := greetingPrefix("Spanish")
		want := "Hola, "
		assertCorrectPrefix(t, got, want)
	})

	t.Run("English prefix", func(t *testing.T) {
		got := greetingPrefix("English")
		want := "Hello, "
		assertCorrectPrefix(t, got, want)
	})

	t.Run("empty string prefix", func(t *testing.T) {
		got := greetingPrefix("")
		want := "Hello, "
		assertCorrectPrefix(t, got, want)
	})

}