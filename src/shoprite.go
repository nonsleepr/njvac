package njvac

import (
	"github.com/antchfx/htmlquery"
)

const ShopRiteURL = "https://covidinfo.reportsonline.com/covidinfo/ShopRite.html"

func GetShopRiteStatus() (string, error) {
	doc, err := htmlquery.LoadURL(ShopRiteURL)
	if err != nil {
		return "", err
	}
	node, err := htmlquery.Query(doc, "/html/body/table/tbody/tr[2]/td/table/tbody/tr/td/div/div/h2/span/span")
	if err != nil {
		return "", err
	}
	return node.FirstChild.Data, nil
}
