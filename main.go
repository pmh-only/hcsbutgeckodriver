package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/firefox"
)

func main() {
	const (
		// These paths will be different on your system.
		seleniumPath    = "bin/selenium-server.jar"
		geckoDriverPath = "bin/geckodriver"
	)

	name := os.Args[1]
	birth := os.Args[2]
	passwd := os.Args[3]
	port := rand.Intn(1024) + 8080

	opts := []selenium.ServiceOption{
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox
	}
	// selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "firefox"}
	caps.AddFirefox(firefox.Capabilities{Args: []string{"-headless"}})
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	if err := wd.Get("https://hcs.eduro.go.kr/"); err != nil {
		panic(err)
	}

	click(wd, "#btnConfirm2")
	click(wd, ".searchBtn")
	click(wd, "#sidolabel")
	click(wd, "#sidolabel>option[value=\"16\"]")
	click(wd, "#crseScCode")
	click(wd, "#crseScCode>option[value=\"4\"]")
	typo(wd, "#orgname", "경북소프트웨어고")
	click(wd, ".searchBtn")
	click(wd, ".layerSchoolArea>li")
	click(wd, ".layerFullBtn")
	typo(wd, "#user_name_input", name)
	typo(wd, "#birthday_input", birth)

	time.Sleep(3 * time.Second)

	click(wd, "#btnConfirm")
	click(wd, "#btnConfirm")
	click(wd, "#btnConfirm")
	click(wd, "#btnConfirm")
	click(wd, "#btnConfirm")

	time.Sleep(3 * time.Second)

	click(wd, "#password")
	click(wd, "#password")
	click(wd, "#password")

	time.Sleep(3 * time.Second)

	for _, c := range passwd {
		click(wd, fmt.Sprintf("a[aria-label=\"%c\"]", c))
	}

	time.Sleep(3 * time.Second)

	click(wd, "#btnConfirm")
	click(wd, "#btnConfirm")
	click(wd, "#btnConfirm")
	click(wd, "#btnConfirm")
	click(wd, "#btnConfirm")

	time.Sleep(3 * time.Second)

	click(wd, ".name")

	time.Sleep(3 * time.Second)

	click(wd, "#survey_q1a1")
	click(wd, "#survey_q2a1")
	click(wd, "#survey_q3a1")
	click(wd, "#survey_q1a1")
	click(wd, "#survey_q2a1")
	click(wd, "#survey_q3a1")
	click(wd, "#survey_q1a1")
	click(wd, "#survey_q2a1")
	click(wd, "#survey_q3a1")

	time.Sleep(3 * time.Second)

	click(wd, "#btnConfirm")
	click(wd, "#btnConfirm")
	click(wd, "#btnConfirm")
	click(wd, "#btnConfirm")
	click(wd, "#btnConfirm")
}

func click(wd selenium.WebDriver, selector string) {
	elem, err := wd.FindElement(selenium.ByCSSSelector, selector)
	if err != nil {
		log.Println(err)
		return
	}

	elem.Click()
}

func typo(wd selenium.WebDriver, selector, text string) {
	elem, err := wd.FindElement(selenium.ByCSSSelector, selector)
	if err != nil {
		log.Println(err)
		return
	}

	elem.SendKeys(text)
}
