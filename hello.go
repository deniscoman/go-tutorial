package main

import "fmt"

const enlgishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return enlgishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Denis"))
}
