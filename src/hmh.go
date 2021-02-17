package njvac

import (
	"github.com/antchfx/htmlquery"
	"strings"
)

const HmhURL = "https://www.hackensackmeridianhealth.org/covid19/covid19-vaccination-scheduling-for-16-64-with-underlying-health-condition/"

func GetHMHStatus() (message string, status bool, err error) {
	doc, err := htmlquery.LoadURL(HmhURL)
	if err != nil {
		return
	}
	node, err := htmlquery.Query(doc, "/html/body/div[1]/div/div/main/section[3]/div/div[1]/div[2]/div/div[1]/strong")
	if err != nil {
		return
	}
	message = node.FirstChild.Data
	status = !strings.Contains(message, "All appointments currently are full")
	return
}
