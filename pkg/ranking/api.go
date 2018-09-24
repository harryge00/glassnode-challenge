package ranking

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RankData struct {
	Name      string `json:"Name"`
	SortOrder string `json:"SortOrder"`
}

type CoinList struct {
	Data map[string]RankData `json:"data"`
}

const coinListURL = "https://www.cryptocompare.com/api/data/coinlist"

func GetCoinList() (*CoinList, error) {
	resp, err := http.Get(coinListURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ticker := CoinList{}
	err = json.Unmarshal(bytes, &ticker)
	return &ticker, nil
}
