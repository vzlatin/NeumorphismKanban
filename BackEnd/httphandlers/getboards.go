package httphandlers

import (
	"encoding/json"
	"log"
	"net/http"

	ws "github.com/vzlatin/NeumorphismKanban/wsconnection"
)

func GetAllBoards(w http.ResponseWriter, r *http.Request) {
	apiConfig, ctx := ws.GetApiConfig()
	boards, err := apiConfig.DB.GetAllBoards(ctx)

	if err != nil {
		log.Printf("error retreiving boards from the database: %s", err)
	}
	response, err := json.Marshal(boards)
	if err != nil {
		log.Printf("error serializing the response data")
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Write(response)
}
