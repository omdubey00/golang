package maps

import "errors"

type Dictionary map[string]string // we have created a type alias for map[string]string

var (
	ErrDictionaryNotFound = errors.New("could not find the word you are looking for") // easier way for
	ErrorWordExists       = errors.New("Word already exists in the dictionary")       // declaration
	ErrorDoesNotExists    = errors.New("Cannot update the word that does not exits")
	ErrorCannotDelete     = errors.New("Cannot delete a word that does not exists")
)

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key] // using the property of map that it returns the value and if that key exists in the map.

	if !ok {
		return "", ErrDictionaryNotFound
	}

	return definition, nil
}

func (d Dictionary) Update(word, newValue string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		{
			d[word] = newValue
		}
	case ErrDictionaryNotFound:
		{
			return ErrorDoesNotExists
		}
	default:
		{
			return err
		}

	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrDictionaryNotFound:
		{
			d[key] = value // since word was not found we add it to the dictionary
		}
	case nil:
		{
			return ErrorWordExists // this is when word is found cuz Search err will be nil if word was found
		}
	default:
		{
			return err
		}
	}
	return nil
}
