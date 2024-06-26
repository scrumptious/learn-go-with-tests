package main

import "fmt"

const englishGreeting = "Hello, "
const spanishGreeting = "Hola, "
const frenchGreeting = "Bonjour, "

type Greeting struct {
	name     string
	language string
}

func Hello(g Greeting) string {
	if g.name == "" {
		g.name = "World"
	}
	if g.language == "" {
		g.language = "English"
	}
	return greetingPrefix(g.language) + g.name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishGreeting
		break
	case "French":
		prefix = frenchGreeting
		break
	default:
		prefix = englishGreeting
	}

	return
}

func main() {
	g1 := Greeting{}
	g2 := Greeting{name: "Lud"}
	fmt.Println(Hello(g1))
	fmt.Println(Hello(g2))
}
