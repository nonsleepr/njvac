package njvac

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const CvsURL = "https://www.cvs.com/immunizations/covid-19-vaccine"
const statusUrl = "https://www.cvs.com/immunizations/covid-19-vaccine.vaccine-status.NJ.json"

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

func GetCVSData() ([]CVSAvailability, string, error) {
	req, err := http.NewRequest("GET", statusUrl, nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Add("Referer", CvsURL)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	var result CVSResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, "", err
	}

	if result.Metadata.Status == "Success" {
		return result.Payload.Data["NJ"], result.Payload.CurrentTime, nil
	} else {
		return nil, "", fmt.Errorf(result.Metadata.Status)
	}
}
