package main

import (
	"github.com/emicklei/go-restful"
	"github.com/harryge00/glassnode-challenge/pkg/ranking"
	"log"
	"net/http"
	"os"
)

func main() {
	ws := new(restful.WebService)
	ws.
		Path("/").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/ranking").To(getRanking))

	restful.Add(ws)
	defaultPort := "8080"
	if os.Getenv("PORT") != "" {
		defaultPort = os.Getenv("PORT")
	}
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))

}

func getRanking(request *restful.Request, response *restful.Response) {
	log.Println("Try to get coinlist.")
	rankingData, err := ranking.GetCoinList()
	if err != nil {
		response.WriteError(http.StatusInternalServerError, err)
	}
	response.WriteAsJson(rankingData)
}
