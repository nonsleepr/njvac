package njvac

import (
	"github.com/antchfx/htmlquery"
)

const ValleyHealthURL = "https://www.valleyhealth.com/covid-19-vaccine-eligibility"

func GetValleyHealthStatus() (string, error) {
	doc, err := htmlquery.LoadURL(ValleyHealthURL)
	if err != nil {
		return "", err
	}
	node, err := htmlquery.Query(doc, "/html/body/div/div/section[3]/section[2]/div/div/div/div/div/div/div/article/div/div/div/p[1]/strong")
	if err != nil {
		return "", err
	}
	return node.FirstChild.Data, nil
}
