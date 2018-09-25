package main

import (
	"github.com/emicklei/go-restful"
	"github.com/harryge00/glassnode-challenge/pkg/price"
	"log"
	"net/http"
)

func main() {
	ws := new(restful.WebService)
	ws.
		Path("/").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/price").To(getPrice).
		Param(restful.QueryParameter("start", "Start ID of coin")).
		Param(restful.QueryParameter("limit", "Size limit of coins")))

	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func getPrice(request *restful.Request, response *restful.Response) {
	start := request.QueryParameter("start")
	limit := request.QueryParameter("limit")
	log.Printf("Try to get coin price. start: %v, limit: %v", start, limit)
	priceMap, err := price.GetPriceMap(start, limit)
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}
	response.WriteAsJson(priceMap)
}
