package helloworld

import "fmt"

type Language struct {
	value string
}

var Languages = struct {
	English Language
	Spanish Language
	French  Language
}{
	English: Language{value: "english"},
	Spanish: Language{value: "spanish"},
	French:  Language{value: "french"},
}

var greetings = map[Language]string{
	Languages.English: "Hello",
	Languages.Spanish: "Hola",
	Languages.French:  "Bonjour",
}

func Hello(name string, language Language) string {
	greeting := greetings[language]

	if name == "" {
		return fmt.Sprintf("%s, world", greeting)
	}

	return fmt.Sprintf("%s, %s", greeting, name)
}
