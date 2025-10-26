package maps

import "testing"

func TestSearch(t *testing.T) {

	t.Run("known word", func(t *testing.T) {
		key := "name"
		value := "theriddler"
		dictionary := Dictionary{key: value}
		got, _ := dictionary.Search(key)
		want := value

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"name": "theriddler"}
		_, err := dictionary.Search("test")
		want := ErrorNotFound

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {

	t.Run("adding a new word", func(t *testing.T) {
		key := "age"
		value := "21"
		dictionary := Dictionary{}
		err := dictionary.Add(key, value)

		assertNoError(t, err)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("adding a existing word", func(t *testing.T) {
		key := "age"
		value := "20"
		dictionary := Dictionary{"age": "20"}
		err := dictionary.Add(key, "21")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("Updating a existing word", func(t *testing.T) {
		key := "name"
		value := "theriddler"
		dictionary := Dictionary{key: value}

		err := dictionary.Update(key, "TheRiddler")
		assertNoError(t, err)
		assertDefinition(t, dictionary, key, "TheRiddler")
	})

	t.Run("Updating a new word", func(t *testing.T) {
		key := "name"
		value := "theriddler"
		dictionary := Dictionary{key: value}

		err := dictionary.Update("age", "21")
		assertError(t, err, ErrCannotUpdate)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestDelete(t *testing.T) {

	t.Run("Deleting an existing word", func(t *testing.T) {
		key := "name"
		value := "theriddler"
		dictionary := Dictionary{key: value}
		dictionary.Delete(key)

		_, error := dictionary.Search(key)
		assertError(t, error, ErrorNotFound)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got : %q , want : %q ", got, want) // why %q in place of %s
	}
}

func assertError(t *testing.T, err, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("Expected an error did not get one")
	}

	if err != want {
		t.Fatal("Did not got the expected error a new error was thrown : ", err)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("Did not expected an error but got one: ", got) // This one is much better because now we get the specific error
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, key, value string) {
	got, err := dictionary.Search(key)
	want := value

	assertNoError(t, err)
	assertStrings(t, got, want)
}
