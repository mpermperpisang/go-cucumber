package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DATA-DOG/godog"
	. "github.com/logrusorgru/aurora"
	"github.com/tebeka/selenium"
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
	btnSuggestion, _ := wd.FindElement(selenium.ByCSSSelector, ".c-btn--large.c-btn--green")

	if err := btnSuggestion.Click(); err != nil {
		wd.Screenshot()
		wd.Quit()
		log.Fatalln(Bold(Red(err)))
	}

	return nil
}

func inputSuggestionForm() error {
	// page object
	fieldName, _ := wd.FindElement(selenium.ByXPATH, "//*[@name='name']")
	radioAge, _ := wd.FindElement(selenium.ByXPATH, "//*[@name='age']")
	optionAge, _ := wd.FindElement(selenium.ByXPATH, "//*[@value='25']")
	fieldOccupation, _ := wd.FindElement(selenium.ByXPATH, "//*[@name='occupation']")
	fieldLocation, _ := wd.FindElement(selenium.ByXPATH, "//*[@name='suggested_location']")
	btnSubmit, _ := wd.FindElement(selenium.ByCSSSelector, ".c-btn--green.u-pad-h--2")

	// action to object
	if err := fieldName.SendKeys("Automation Dweb"); err != nil {
		wd.Screenshot()
		wd.Quit()
		log.Fatalln(Bold(Red(err)))
	}
	if err := radioAge.Click(); err != nil {
		wd.Screenshot()
		wd.Quit()
		log.Fatalln(Bold(Red(err)))
	}
	if err := optionAge.Click(); err != nil {
		wd.Screenshot()
		wd.Quit()
		log.Fatalln(Bold(Red(err)))
	}
	if err := fieldOccupation.SendKeys("SDET"); err != nil {
		wd.Screenshot()
		wd.Quit()
		log.Fatalln(Bold(Red(err)))
	}
	if err := fieldLocation.SendKeys("Tangerang"); err != nil {
		wd.Screenshot()
		wd.Quit()
		log.Fatalln(Bold(Red(err)))
	}
	if err := btnSubmit.Click(); err != nil {
		wd.Screenshot()
		wd.Quit()
		log.Fatalln(Bold(Red(err)))
	}

	return nil
}

func validateSuggestionInputSuccess(msg string) error {
	submitSuccess, _ := wd.FindElement(selenium.ByXPATH, "//*[contains(text(),'"+os.Getenv("VALIDATION")+"')]")

	if _, err := submitSuccess.IsDisplayed(); err != nil {
		wd.Screenshot()
		wd.Quit()
		log.Println(Bold(Red("TIDAK KETEMU")))
	}

	defer wd.Quit()
	return nil
}

func BukaBikeLandingPageSteps(s *godog.Suite) {
	s.Step(`^user access bike url$`, accessBukaBikeLandingPage)
	s.Step(`^user fill suggestion form$`, suggestionForm)
	s.Step(`^user must get success message "([^\"]*)"$`, validateSuggestionInputSuccess)

	s.BeforeScenario(func(interface{}) {
		Godogs = 0 // clean the state before every scenario
	})
}
