package main

import "testing" // this is my testing framework

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got : %q , want : %q ", got, want)
		}
	}

	t.Run("says hello to anyone", func(t *testing.T) {
		got := Hello("TheRiddler", "")
		want := "Hello TheRiddler"
		assertCorrectMessage(t, got, want)
	})

	t.Run("says hello darling when empty string is given in hello fn", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola darling"
		assertCorrectMessage(t, got, want)
	})
}
