package main

import (
	"fmt"
	"github.com/nonsleepr/njvac/src"
	"github.com/tomlazar/table"
	"os"
)

const zipCode = "07450"

func statusString(status bool) string {
	if status {
		return "✅"
	}
	return "⛔"
}

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
	hmh, status, err := njvac.GetHMHStatus()
	if err == nil {
		fmt.Println(statusString(status), hmh)
	} else {
		fmt.Println("Failed to retrieve Hackensack Meridian Health data:", err)
	}

	fmt.Printf("\n\n==== ShopRite (%s) ====\n\n", njvac.ShopRiteURL)
	shopRite, status, err := njvac.GetShopRiteStatus()
	if err == nil {
		fmt.Println(statusString(status), shopRite)
	} else {
		fmt.Println("Failed to retrieve ShopRite data:", err)
	}

	fmt.Printf("\n\n==== Valley Health (%s) ====\n\n", njvac.ValleyHealthURL)
	valleyHealth, status, err := njvac.GetValleyHealthStatus()
	if err == nil {
		fmt.Println(statusString(status), valleyHealth)
	} else {
		fmt.Println("Failed to retrieve Valley Health data:", err)
	}

	fmt.Printf("\n\n==== RiteAid (%s) ====\n\n", njvac.RiteAidURL)
	riteaidData, err := njvac.GetRiteAidStatus(zipCode)
	if err == nil {
		table.MarshalTo(os.Stdout, riteaidData, &table.Config{})
	} else {
		fmt.Println("Failed to retrieve RiteAid data:", err)
	}

}
