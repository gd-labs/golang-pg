package hello

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHello = "Hello, "
	spanishHello = "Hola, "
	frenchHello  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("%s%s!", getPrefix(language), name)
}

func getPrefix(language string) string {
	switch language {
	case spanish:
		return spanishHello
	case french:
		return frenchHello
	default:
		return englishHello
	}
}
