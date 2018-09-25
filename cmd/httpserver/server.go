package main

import (
	"encoding/json"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/harryge00/glassnode-challenge/pkg/price"
	"github.com/harryge00/glassnode-challenge/pkg/ranking"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	rankURL  string
	priceURL string
)

type Coin struct {
	Symbol string  `json:"symbol"`
	Rank   int     `json:"rank"`
	Price  float64 `json:"price,omitempty"`
}

func main() {
	ws := new(restful.WebService)
	ws.
		Path("/").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/").To(topCoinList).
		Param(restful.QueryParameter("limit", "Size limit of coins")))

	restful.Add(ws)

	rankURL = fmt.Sprintf("http://%s:%s/ranking", os.Getenv("RANK_ADDR"), os.Getenv("RANK_PORT"))
	priceURL = fmt.Sprintf("http://%s:%s/price", os.Getenv("PRICE_ADDR"), os.Getenv("PRICE_PORT"))
	log.Fatal(http.ListenAndServe(":6667", nil))
}

func topCoinList(request *restful.Request, response *restful.Response) {
	log.Println("Try to get coinlist.")
	var err error
	limit := 100
	limitQuery := request.QueryParameter("limit")
	if limitQuery != "" {
		limit, err = strconv.Atoi(limitQuery)
		if err != nil {
			response.WriteError(http.StatusBadRequest, err)
			return
		}
	}
	log.Println(limit)
	resp, err := http.Get(rankURL)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	var coinList ranking.CoinList
	err = json.Unmarshal(bytes, &coinList)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
		return
	}
	log.Println(coinList)

	topCoins := make([]Coin, limit)
	for _, coin := range coinList.Data {
		rank, _ := strconv.Atoi(coin.SortOrder)
		if rank <= limit {
			topCoins[rank-1].Symbol = coin.Name
			topCoins[rank-1].Rank = rank
		}
	}

	// Get prices for 100 coins every time
	for i := 0; i < limit; i += 100 {
		var priceList price.PriceMap
		url := fmt.Sprintf("%v?start=%v", priceURL, i)
		log.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			response.WriteError(http.StatusInternalServerError, err)
			return
		}
		defer resp.Body.Close()
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			response.WriteError(http.StatusInternalServerError, err)
			return
		}

		err = json.Unmarshal(bytes, &priceList)
		if err != nil {
			response.WriteError(http.StatusInternalServerError, err)
			return
		}
		log.Println(priceList.PriceMap)
		for i := range topCoins {
			log.Println(topCoins[i].Symbol, priceList.PriceMap[topCoins[i].Symbol])
			if coinPrice, exist := priceList.PriceMap[topCoins[i].Symbol]; exist {
				topCoins[i].Price = coinPrice
			}
		}
	}
	response.WriteAsJson(topCoins)
}
