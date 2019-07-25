package main

import (
	"log"

	"github.com/DATA-DOG/godog"
)

var usersName, meetName string

func givenName(name string) error {
	usersName = name

	return nil
}

func meetUser() error {
	meetName = hello(usersName)

	return nil
}

func sayHello(greet string) error {
	if meetName != greet {
		log.Fatalf("hello(\"\") failed, expected %v, got %v", "Hello Dude!", meetName)
	}

	return nil
}

func UnitTestCucumber(s *godog.Suite) {
	s.Step(`^user has a name "([^\"]*)"$`, givenName)
	s.Step(`^neighbor meet user$`, meetUser)
	s.Step(`^neighbor say "([^\"]*)"$`, sayHello)
}
