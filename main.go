package main

import (
	"fmt"
	"github.com/nonsleepr/go-nj-vaccine/src"
	"github.com/tomlazar/table"
	"os"
)

func main() {
	cvsData, cvsTimestamp, err := njvac.GetCVSData()
	if err != nil {
		return
	}
	fmt.Printf("Data as of %s\n", cvsTimestamp)
	table.MarshalTo(os.Stdout, cvsData, &table.Config{})
}
