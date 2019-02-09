package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/marcopaganini/gosmart"
	"github.com/mitchellh/go-homedir"
	autohouseClient "github.com/solcates/autohouse/client"
	"io/ioutil"
	"path/filepath"
)

var c autohouseClient.Client
var config = &Config{}
var hd, _ = homedir.Dir()
var configFile = filepath.Join(hd, ".autohouse.config")
var tokenFile = filepath.Join(hd, ".autohouse.token")

type Config struct {
	AppClientID     string `json:"app_client_id"`
	AppClientSecret string `json:"app_client_secret"`
}

func init() {

	// See if we get a config
	configData, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configData, config)
	if err != nil {
		panic(err)
	}
}


func main() {

	oConfig := gosmart.NewOAuthConfig(config.AppClientID, config.AppClientSecret)
	token, err := gosmart.GetToken(tokenFile, oConfig)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	client := oConfig.Client(ctx, token)
	endpoint, err := gosmart.GetEndPointsURI(client)
	if err != nil {
		panic(err)
	}

	c = &autohouseClient.AutoHouse{
		Client:   client,
		Endpoint: endpoint,
	}

	locations, err := c.Locations()
	if err != nil {
		panic(err)
	}
	fmt.Println(locations.Hubs)
}
