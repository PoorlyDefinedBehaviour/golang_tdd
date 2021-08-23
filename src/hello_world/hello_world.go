package helloworld

import "fmt"

type Language struct {
	language string
}

var Languages = struct {
	English Language
	Spanish Language
}{
	English: Language{language: "english"},
	Spanish: Language{language: "spanish"},
}

var greetings = map[Language]string{
	Languages.English: "Hello",
	Languages.Spanish: "Hola",
}

func greetingByLanguage(language Language) string {
	greeting, ok := greetings[language]
	if !ok {
		return greetings[Languages.English]
	}

	return greeting
}

func Hello(name string, language Language) string {
	greeting := greetingByLanguage(language)

	if name == "" {
		return fmt.Sprintf("%s, world", greeting)
	}

	return fmt.Sprintf("%s, %s", greeting, name)
}
