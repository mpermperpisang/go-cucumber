package main

import (
	"fmt"

	"github.com/DATA-DOG/godog"
	"github.com/tebeka/selenium"
)

func openBukalapak() error {
	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	if err := wd.Get("localhost:8080/?name=mpermperpisang"); err != nil {
		panic(err)
	}

	fmt.Println(wd.CurrentURL())
	
	currentTitle, _ := wd.Title()
	if expectTitle := ("Welcome mpermperpisang"); currentTitle != expectTitle {
		panic("Nama tidak sesuai")
	}

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
	s.Step(`^open browser$`, openBukalapak)

	s.BeforeScenario(func(interface{}) {
		Godogs = 0 // clean the state before every scenario
	})
}
