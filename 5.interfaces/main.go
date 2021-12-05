package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Print(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// very custom logic for generating english greeting
	return "Hi There!"
}

func (eb spanishBot) getGreeting() string {
	// very custom logic for generating english greeting
	return "Hi There!"
}
