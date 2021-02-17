package njvac

import (
	"github.com/antchfx/htmlquery"
	"strings"
)

const ValleyHealthURL = "https://www.valleyhealth.com/covid-19-vaccine-eligibility"

func GetValleyHealthStatus() (message string, status bool, err error) {
	doc, err := htmlquery.LoadURL(ValleyHealthURL)
	if err != nil {
		return
	}
	node, err := htmlquery.Query(doc, "/html/body/div/div/section[3]/section[2]/div/div/div/div/div/div/div/article/div/div/div/p[1]/strong")
	if err != nil {
		return
	}
	message = node.FirstChild.Data
	status = !strings.Contains(message, "NO APPOINTMENTS ARE AVAILABLE")
	return
}
