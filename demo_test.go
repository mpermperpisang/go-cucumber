package main

import (
	"fmt"
	"os"

	"github.com/DATA-DOG/godog"
	. "github.com/logrusorgru/aurora"
)

var url string

func websiteURL() error {
	url = os.Getenv("URL_WEBSITE")

	return nil
}

func accessWebsiteURL() error {
	seleniumWebDriverConnect()
	GetURL(os.Getenv("URL_WEBSITE"))

	return nil
}

func validateWebsiteURL() error {
	if expectURL := (url); url != expectURL {
		wd.Screenshot()
		fmt.Println(Bold(Red("URL DOESN'T MATCH")))
	} else {
		fmt.Println(Bold(Magenta("SCENARIO IS SUCCESS")))
	}

	return nil
}

func GodogDemoSteps(s *godog.Suite) {
	s.Step(`^user has url$`, websiteURL)
	s.Step(`^user access the url$`, accessWebsiteURL)
	s.Step(`^user must directed to correct url$`, validateWebsiteURL)

	s.BeforeScenario(func(interface{}) {
		Godogs = 0 // clean the state before every scenario
	})

	s.AfterScenario(func(interface{}, error) {
		wd.Quit()
		Godogs = 0 // clean the state before every scenario
	})
}
