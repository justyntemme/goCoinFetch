package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func round(n string) string {
	f, err := strconv.ParseFloat(n, 32)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%.2f", f)
}

func GrabTicker(coinHandle string) string {
	Ccoin := []coin{}

	resp, err := http.Get("https://api.nexchange.io/en/api/v1/price/" + coinHandle + "USD/latest?format=json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(resp.Body).Decode(&Ccoin)
	if err != nil {
		log.Fatal(err)
	}

	ask := round(Ccoin[0].Cticker.Ask)
	bid := round(Ccoin[0].Cticker.Bid)

	return fmt.Sprintf("Ask: %s Bid: %s", ask, bid)
}
