package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-cucumber/features/object_abstractions"

	"github.com/DATA-DOG/godog"
	. "github.com/logrusorgru/aurora"
)

func accessBukaBikeLandingPage() error {
	seleniumWebDriverConnect()

	if err := wd.Get(baseUrl + os.Getenv("URL_BIKE")); err != nil {
		log.Fatalln(Bold(Red(err)))
	}

	fmt.Println(wd.CurrentURL())

	return nil
}

func suggestionForm() error {
	clickYesIWantButton()
	inputSuggestionForm()

	return nil
}

func clickYesIWantButton() error {
	FindElementByCss(object_abstractions.BtnSuggestion).Click()
	return nil
}

func inputSuggestionForm() error {
	FindElementByXpath(object_abstractions.FieldName).SendKeys("Automation Dweb")
	FindElementByXpath(object_abstractions.RadioAge).Click()
	FindElementByXpath(object_abstractions.OptionAge).Click()
	FindElementByXpath(object_abstractions.FieldOccupation).SendKeys("SDET")
	FindElementByXpath(object_abstractions.FieldLocation).SendKeys("Tangerang")
	FindElementByCss(object_abstractions.BtnSubmit).Click()

	return nil
}

func validateSuggestionInputSuccess(msg string) error {
	if _, err := FindElementByXpath("//*[contains(text(),'" + os.Getenv("VALIDATION") + "')]").IsDisplayed(); err != nil {
		wd.Screenshot()
		log.Println(Bold(Red("TIDAK KETEMU")))
	}

	wd.Quit()
	return nil
}

func BukaBikeLandingPageSteps(s *godog.Suite) {
	s.Step(`^user access bike url$`, accessBukaBikeLandingPage)
	s.Step(`^user fill suggestion form$`, suggestionForm)
	s.Step(`^user must not see window "([^\"]*)"$`, validateSuggestionInputSuccess)

	s.BeforeScenario(func(interface{}) {
		Godogs = 0 // clean the state before every scenario
	})
}
