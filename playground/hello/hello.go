package main

import (
	"anve/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Anve")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
