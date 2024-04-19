package main

import "fmt"

func main() {
	ninja1, ninja2 := make(chan string), make(chan string)
	go electNinja(ninja1, "Ninja 1")
	go electNinja(ninja2, "Ninja 2")

	select {
	case message := <-ninja1:
		fmt.Print(message)

	case message := <-ninja2:
		fmt.Print(message )
	}
}

func electNinja(ninja chan string, message string) {
	ninja <- message
}
