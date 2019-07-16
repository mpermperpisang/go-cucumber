package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	. "github.com/logrusorgru/aurora"
	"github.com/tebeka/selenium"
)

var baseUrl string
var Godogs int
var wd selenium.WebDriver

func init() {
	env := godotenv.Load()
	if env != nil {
		log.Println(Bold(Red(env)))
	}

	baseUrl = os.Getenv("URL")
}

func main() {
}

func seleniumWebDriverConnect() error {
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, _ = selenium.NewRemote(caps, "")

	return nil
}
