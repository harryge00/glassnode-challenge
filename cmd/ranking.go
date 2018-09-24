package main

import (
	"net/http"
	"github.com/emicklei/go-restful"
	"github.com/harryge00/glassnode-challenge/pkg/ranking"
	"log"
)

func main() {
	ws := new(restful.WebService)
	ws.
		Path("/").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/ranking").To(getRanking))

	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func getRanking(request *restful.Request, response *restful.Response) {
	log.Println("Try to get coinlist.")
	rankingData, err := ranking.GetCoinList()
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}
	response.WriteAsJson(rankingData)
}