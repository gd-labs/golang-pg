package main

import "fmt"

const englishHello = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("%s%s!", englishHello, name)
}
