package main

import "fmt"

const englishGreeting = "Hello, "

func Hello(n string) string {
	if n == "" {
		n = "World"
	}
	return englishGreeting + n
}

func main() {
	fmt.Println(Hello(""))
	fmt.Println(Hello("Lud"))
}
