package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word ", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("unknown word ", func(t *testing.T) {

		_, err := dictionary.Search("unknown")
		want := ErrDictionaryNotFound.Error()

		if err == nil {
			t.Fatal("expected an error but didn't get one ")
		}
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	t.Run("Adding a word to the dictionary", func(t *testing.T) {
		dictionary.Add("theriddler", "om")
		want := "om"
		got, err := dictionary.Search("theriddler")
		assertNoError(t, err)
		assertString(t, got, want)
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got : %s , want : %s ", got, want)
	}
}

func assertError(t *testing.T, got error, want string) {
	t.Helper()
	if got.Error() != want {
		t.Errorf("got : %q , want : %q ", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	if err != nil {
		t.Error("Did not expected an error but got one ")
	}
}
