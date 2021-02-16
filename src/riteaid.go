package njvac

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const RiteAidURL = "https://www.riteaid.com/pharmacy/covid-qualifier"

const riteAidStores = "https://www.riteaid.com/services/ext/v2/stores/getStores?attrFilter=PREF-112&fetchMechanismVersion=2&radius=50"

const checkSlots = "https://www.riteaid.com/services/ext/v2/vaccine/checkSlots"

type RiteAidSlots struct {
	Address string `table:"STORE ADDRESS"`
	Slot1   bool   `json:"1" table:"SLOT 1 AVAILABILITY"`
	Slot2   bool   `json:"2" table:"SLOT 2 AVAILABILITY"`
}

type RiteAidStore struct {
	StoreNumber int
	Address     string
	City        string
	State       string
	ZipCode     string
}

func (store *RiteAidStore) String() string {
	return fmt.Sprintf("%s, %s, %s, %s", store.Address, store.City, store.State, store.ZipCode)
}

func getStoreSlots(storeNumber int) (slots RiteAidSlots, err error) {
	checkSlotsURL, err := url.Parse(checkSlots)
	if err != nil {
		return
	}
	query := checkSlotsURL.Query()
	query.Add("storeNumber", strconv.Itoa(storeNumber))
	checkSlotsURL.RawQuery = query.Encode()

	resp, err := http.Get(checkSlotsURL.String())
	if err != nil {
		return RiteAidSlots{}, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var result struct {
		Status string
		Data   struct {
			Slots RiteAidSlots
		}
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return
	}
	if result.Status == "SUCCESS" {
		return result.Data.Slots, nil
	} else {
		err = fmt.Errorf(result.Status)
		return
	}
}

func GetRiteAidStatus(zip string) ([]RiteAidSlots, error) {
	riteAidURL, err := url.Parse(riteAidStores)
	if err != nil {
		return nil, err
	}
	query := riteAidURL.Query()
	query.Add("address", zip)
	riteAidURL.RawQuery = query.Encode()

	resp, err := http.Get(riteAidURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result struct {
		Status string
		Data   struct {
			Stores []RiteAidStore
		}
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	if result.Status != "SUCCESS" {
		return nil, fmt.Errorf(result.Status)
	}

	storeSlots := make([]RiteAidSlots, 0, len(result.Data.Stores))
	for _, store := range result.Data.Stores {
		slots, err := getStoreSlots(store.StoreNumber)
		slots.Address = store.String()
		storeSlots = append(storeSlots, slots)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to get store slots for", store.StoreNumber, err)
		}
	}
	return storeSlots, nil
}
