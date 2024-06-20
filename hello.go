package main

import "fmt"

const englishGreeting = "Hello, "
const spanishGreeting = "Hola, "

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

	switch g.language {
	case "Spanish":
		return spanishGreeting + g.name
	}
	return englishGreeting + g.name
}

func main() {
	g1 := Greeting{}
	g2 := Greeting{name: "Lud"}
	fmt.Println(Hello(g1))
	fmt.Println(Hello(g2))
}
