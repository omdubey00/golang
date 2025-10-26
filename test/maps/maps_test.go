package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"name": "theriddler"}
	got := dictionary.Search("name")
	want := "theriddler"

	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got : %q , want : %q ", got, want) // why %q in place of %s
	}
}
