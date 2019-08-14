package main

import (
	"sourcegraph.com/sourcegraph/go-selenium"
	"testing"
)

var caps selenium.Capabilities
var executorURL = "http://localhost:4444/wd/hub"

// An example test using the WebDriverT and WebElementT interfaces. If you use the non-*T
// interfaces, you must perform error checking that is tangential to what you are testing,
// and you have to destructure results from method calls.
func TestWithT(t *testing.T) {
	wd, _ := selenium.NewRemote(caps, executor)

	// Call .T(t) to obtain a WebDriverT from a WebDriver (or to obtain a WebElementT from
	// a WebElement).
	wdt := wd.T(t)

	// Calls `t.Fatalf("Get: %s", err)` upon failure.
	wdt.Get("http://example.com")

	// Calls `t.Fatalf("FindElement(by=%q, value=%q): %s", by, value, err)` upon failure.
	elem := wdt.FindElement(selenium.ByCSSSelector, ".foo")

	// Calls `t.Fatalf("Text: %s", err)` if the `.Text()` call fails.
	if elem.Text() != "bar" {
		t.Fatalf("want elem text %q, got %q", "bar", elem.Text())
	}
}
