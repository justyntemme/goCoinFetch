package web

import (
	"encoding/json"
	"log"
	"net/http"
)

type ticker struct {
	Ask string `json:"ask"`
	Bid string `json:"bid"`
}

type coin struct {
	CreatedOn string `json:"created_on"`
	UnixTime  string `json:"unix_time"`
	Cticker   ticker `json:"ticker"`
}

func GrabTicker(coinHandle string) string {
	price := ""

	Ccoin := []coin{}

	resp, err := http.Get("https://api.nexchange.io/en/api/v1/price/" + coinHandle + "USD/latest?format=json")
	if err != nil {

		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(&Ccoin)
	if err != nil {

		log.Fatal(err)
	}

	price = "Ask: " + Ccoin[0].Cticker.Ask + " Bid: " + Ccoin[0].Cticker.Bid
	return price
}
