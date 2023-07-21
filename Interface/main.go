package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishhBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishhBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hi There!"
}

func (eb spanishhBot) getGreeting() string {
	return "Hola!"
}
