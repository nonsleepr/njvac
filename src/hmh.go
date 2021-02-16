package njvac

import (
	"github.com/antchfx/htmlquery"
)

const url = "https://www.hackensackmeridianhealth.org/covid19/covid19-vaccination-scheduling-for-16-64-with-underlying-health-condition/"

func GetHMHStatus() (string, error) {
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		return "", err
	}
	node, err := htmlquery.Query(doc, "/html/body/div[1]/div/div/main/section[3]/div/div[1]/div[2]/div/div[1]/strong")
	if err != nil {
		return "", err
	}
	return node.FirstChild.Data, nil
}
