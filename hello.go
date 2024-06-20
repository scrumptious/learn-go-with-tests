package main

import "fmt"

func Hello(n string) string {
	return fmt.Sprintf("Hello %s", n)
}

func main() {
	fmt.Println(Hello("Sam"))
}
