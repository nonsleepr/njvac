package main

import (
	"fmt"
	"github.com/nonsleepr/njvac/src"
	"github.com/tomlazar/table"
	"os"
)

func main() {
	fmt.Printf("\n\n==== CVS Vaccination (%s) ====\n\n", njvac.CvsURL)
	cvsData, cvsTimestamp, err := njvac.GetCVSData()
	if err == nil {
		fmt.Printf("Data as of %s\n\n", cvsTimestamp)
		table.MarshalTo(os.Stdout, cvsData, &table.Config{})
	} else {
		fmt.Println("Failed to retrieve CVS data:", err)
	}

	fmt.Printf("\n\n==== Hackensack Meridian Health (%s) ====\n\n", njvac.HmhURL)
	hmh, err := njvac.GetHMHStatus()
	if err == nil {
		fmt.Println(hmh)
	} else {
		fmt.Println("Failed to retrieve Hackensack Meridian Health data:", err)
	}

	fmt.Printf("\n\n==== ShopRite (%s) ====\n\n", njvac.ShopRiteURL)
	shopRite, err := njvac.GetShopRiteStatus()
	if err == nil {
		fmt.Println(shopRite)
	} else {
		fmt.Println("Failed to retrieve ShopRite data:", err)
	}

	fmt.Printf("\n\n==== RiteAid (%s) ====\n\n", njvac.RiteAidURL)
	riteaidData, err := njvac.GetRiteAidStatus("07450")
	if err == nil {
		table.MarshalTo(os.Stdout, riteaidData, &table.Config{})
	} else {
		fmt.Println("Failed to retrieve RiteAid data:", err)
	}

}
