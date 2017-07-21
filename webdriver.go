package webdriver

import (
	"time"

	"github.com/tebeka/selenium"
)

// SetUpChrome sets the chromedriver
func SetUpChrome() selenium.WebDriver {
	caps := selenium.Capabilities{"browserName": "chrome"}
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}

	// Setup wait times to 30 secs
	driver.SetImplicitWaitTimeout(time.Millisecond * 30000)
	driver.SetAsyncScriptTimeout(time.Millisecond * 30000)
	driver.SetPageLoadTimeout(time.Millisecond * 30000)

	return driver
}
