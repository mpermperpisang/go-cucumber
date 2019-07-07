package main

import (
	"fmt"
	"log"

	"github.com/DATA-DOG/godog"
	. "github.com/logrusorgru/aurora"
	"github.com/tebeka/selenium"
)

var wd selenium.WebDriver

func seleniumWebDriverConnect() error {
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, _ = selenium.NewRemote(caps, "")

	return nil
}

func accessURL(url string) error {
	if err := wd.Get(fmt.Sprintf("localhost:8080/?name=%s", url)); err != nil {
		panic(err)
	}

	fmt.Println(wd.CurrentURL())

	return nil
}

func validateWindowTitle(name string) error {
	currentTitle, _ := wd.Title()
	if expectTitle := (fmt.Sprintf("Welcome %s", name)); currentTitle != expectTitle {
		wd.Screenshot()
		wd.Quit()
		log.Fatalln(Bold(Red("NAMA TIDAK SESUAI")))
	} else {
		fmt.Println(Bold(Magenta("SKENARIO SUKSES")))
	}

	defer wd.Quit()
	return nil
}

func thereAreGodogs(available int) error {
	Godogs = available
	return nil
}

func iEat(num int) error {
	if Godogs < num {
		return fmt.Errorf("you cannot eat %d godogs, there are %d available", num, Godogs)
	}
	Godogs -= num
	return nil
}

func thereShouldBeRemaining(remaining int) error {
	if Godogs != remaining {
		return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, Godogs)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there are (\d+) godogs$`, thereAreGodogs)
	s.Step(`^I eat (\d+)$`, iEat)
	s.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)
	s.Step(`^user open browser$`, seleniumWebDriverConnect)
	s.Step(`^user access url with name "([^\"]*)"$`, accessURL)
	s.Step(`^user must get window title Welcome "([^\"]*)"$`, validateWindowTitle)

	s.BeforeScenario(func(interface{}) {
		Godogs = 0 // clean the state before every scenario
	})
}
