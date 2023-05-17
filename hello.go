package main

import (
	"fmt"
)

const (
	spanish = "Spanish"
	french  = "French"

	enlgishHelloPrefix = "Hello, "
	spansihHelloPrefix = "!Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spansihHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = enlgishHelloPrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello("Denis", spanish))
}
