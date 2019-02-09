package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client interface {
	Locations() (locations Location, err error)
}

type AutoHouse struct {
	Client   *http.Client
	Endpoint string
}

func (ah *AutoHouse) Locations() (locations Location, err error) {

	var resp *http.Response
	var data []byte
	resp, err = ah.Client.Get(ah.Endpoint + "/locations")
	if err != nil {
		return
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(data, &locations)

	return

}
