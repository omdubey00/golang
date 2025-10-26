package maps

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrorNotFound   = DictionaryErr("could not find the word you are looking for")
	ErrWordExists   = DictionaryErr("Duplicate key ,Cannot Add the word")
	ErrCannotUpdate = DictionaryErr("Cannot update the word , as it does not exists in the dictionary")
	ErrCannotDelete = DictionaryErr("Cannot delete the word ,as it does not exists in the dictionary")
)

func (d Dictionary) Search(key string) (value string, err error) {
	value, found := d[key]
	if !found {
		return "", ErrorNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) (err error) {
	_, ok := d[key]
	if ok {
		return ErrWordExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) (err error) {
	_, ok := d[key]
	if !ok {
		return ErrCannotUpdate
	}
	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
