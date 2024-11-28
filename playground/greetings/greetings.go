package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("name can not be empty")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Hello, %v. Welcome!",
		"Hola, %v. Welcome!",
		"Bonjour, %v. Welcome!",
		"Ciao, %v. Welcome!",
		"Hallo, %v. Welcome!",
		"Ciao, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"Welcome, %v!",
	}

	return formats[rand.Intn(len(formats))]
}
