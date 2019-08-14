package main

import (
	"fmt"
	"regexp"
	"sourcegraph.com/sourcegraph/go-selenium"
	"time"
)

func main() {
	var webDriver selenium.WebDriver
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	if webDriver, err = selenium.NewRemote(caps, "http://localhost:8081/wd/hub"); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}
	defer webDriver.Quit()

	err = webDriver.Get("http://ttvn.vn/")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

	webDriver.ExecuteScript("window.scrollTo(0, document.body.scrollHeight)", []interface{}{});

	time.Sleep(4 * time.Second)
	if title, err := webDriver.Title(); err == nil {
		fmt.Printf("Page title: %s\n", title)
	} else {
		fmt.Printf("Failed to get page title: %s", err)
		return
	}

	var elems []selenium.WebElement


	var re = regexp.MustCompile(`.*\/[a-zA-Z0-9\-]+\/[a-zA-Z0-9\-]+\d{15,}\.htm$`)

	elems, err = webDriver.FindElements(selenium.ByTagName, "a")
	if err != nil {
		fmt.Printf("Failed to find element: %s\n", err)
		return
	}

	var hrefs []string
	fmt.Printf("Number of elements: %d\n", len(elems))

	flag := 0
	for _, elem := range elems {
		//text, _ := elem.Text()
		//fmt.Println("Text" + text)
		href, err := elem.GetAttribute("href");

		if err == nil {
			fmt.Printf("Repository: %s\n", href)
			if re.MatchString(href){
				flag = 0
				for _, link := range hrefs{
					if link == href{
						flag = 1
						break
					}
				}
				if flag == 0{
					hrefs = append(hrefs, href)
				}

			}
		} else {
			fmt.Printf("Failed to get href of element: %s\n", err)
			return
		}
	}
	fmt.Println(len(hrefs))
	for _, link := range hrefs {
		fmt.Println(link)
	}

	// output:
	// Page title: go-selenium - Sourcegraph
	// Repository: go-selenium
}