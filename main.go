package main

import (
	"encoding/json"
	"fmt"
	"github.com/tomlazar/table"
	"io/ioutil"
	"net/http"
	"os"
)

const REFERER = "https://www.cvs.com/immunizations/covid-19-vaccine"
const STATUS_URL = "https://www.cvs.com/immunizations/covid-19-vaccine.vaccine-status.NJ.json"

type CVSAvailability struct {
	Available    int    `table:"TOTAL AVAILABLE" json:"totalAvailable,string"`
	City         string `table:"CITY"`
	State        string `table:"STATE"`
	PctAvailable string `table:"PERCENT AVAILABLE"`
	Status       string `table:"STATUS"`
}

type CVSResult struct {
	Payload struct {
		CurrentTime string
		Data        map[string][]CVSAvailability
	} `json:"responsePayloadData"`
	Metadata struct {
		Status string `json:"statusDesc"`
	} `json:"responseMetaData"`
}

func main() {
	req, _ := http.NewRequest("GET", STATUS_URL, nil)
	req.Header.Add("Referer", REFERER)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	var result CVSResult
	json.Unmarshal(data, &result)

	if result.Metadata.Status == "Success" {
		fmt.Printf("Data as of %s\n", result.Payload.CurrentTime)
		table.MarshalTo(os.Stdout, result.Payload.Data["NJ"], &table.Config{})
	}
}
