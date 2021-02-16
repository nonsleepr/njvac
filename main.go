package main

import (
	"fmt"
	"github.com/nonsleepr/njvac/src"
	"github.com/tomlazar/table"
	"os"
)

func main() {
	fmt.Println("\n\n==== CVS Vaccination ====\n")
	cvsData, cvsTimestamp, err := njvac.GetCVSData()
	if err == nil {
		fmt.Printf("Data as of %s\n\n", cvsTimestamp)
		table.MarshalTo(os.Stdout, cvsData, &table.Config{})
	} else {
		fmt.Println("Failed to retrieve CVS data:", err)
	}

	fmt.Println("\n\n==== Hackensack Meridian Health ====\n")
	hmh, err := njvac.GetHMHStatus()
	if err == nil {
		fmt.Println(hmh)
	} else {
		fmt.Println("Failed to retrieve Hackensack Meridian Health data:", err)
	}
}
