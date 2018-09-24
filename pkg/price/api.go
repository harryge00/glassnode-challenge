package price

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Ticker struct {
	Data map[string]TickerData `json:"data"`
}

type Quote struct {
	Price float64 `json:"price"`
}

type TickerData struct {
	Symbol string           `json:"symbol"`
	Quotes map[string]Quote `json:"quotes"`
}

type PriceMap struct {
	PriceMap map[string]float64 `json:"pricemap"`
}

func GetPriceMap(start string, limit string) (*PriceMap, error) {
	url := fmt.Sprintf("https://api.coinmarketcap.com/v2/ticker/?start=%s&limit=%s", start, limit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ticker := Ticker{}
	err = json.Unmarshal(bytes, &ticker)
	if err != nil {
		return nil, err
	}
	var result PriceMap
	result.PriceMap = make(map[string]float64)
	for _, data := range ticker.Data {
		result.PriceMap[data.Symbol] = data.Quotes["USD"].Price
	}

	return &result, nil
}
