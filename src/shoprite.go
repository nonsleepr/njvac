package njvac

import (
	"github.com/antchfx/htmlquery"
	"strings"
)

const ShopRiteURL = "https://covidinfo.reportsonline.com/covidinfo/ShopRite.html"

func GetShopRiteStatus() (message string, status bool, err error) {
	doc, err := htmlquery.LoadURL(ShopRiteURL)
	if err != nil {
		return
	}
	node, err := htmlquery.Query(doc, "/html/body/table/tbody/tr[2]/td/table/tbody/tr/td/div/div/h2/span/span")
	if err != nil {
		return
	}
	message = node.FirstChild.Data
	status = !strings.Contains(message, "There are currently no COVID-19 vaccine appointments")
	return
}
