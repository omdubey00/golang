package main

import _ "fmt"

func Hello(name string, language string) string { // Public function
	if name == "" {
		name = "darling"
	}
	return grettingLanguage(language) + name
}

func grettingLanguage(language string) (prefix string) { // Private function
	switch language {
	case "Spanish":
		prefix = "Hoal"
	case "French":
		prefix = "Bonjour"
	default:
		prefix = "Hello"
	}
	return
}
