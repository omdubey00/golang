package maps

type Dictionary map[string]string

func (d Dictionary) Search(key string) (value string) {
	value = d[key]
	return
	// because this func is not going to interact or change anything in the map jsut going to look through the map so searching through the real one or the copied one hence pointer to it does not matter so that is why the choice for pointer or not.
}
