package httphandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	ws "github.com/vzlatin/NeumorphismKanban/wsconnection"
)

func GetBoardIdData(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var uuidString string
	err := decoder.Decode(&uuidString)
	if err != nil {
		log.Printf("error: decoding the request body: %s", err)
	}

	uuid, err := uuid.Parse(uuidString)
	if err != nil {
		log.Printf("error: couldn't parse the incoming uuid string: %s", err)
	}
	apiConfig, ctx := ws.GetApiConfig()
	boardData, err := apiConfig.DB.GetBoardIdData(ctx, uuid)
	if err != nil {
		log.Printf("error retrieving board id data: %s", err)
	}

	response, err := json.Marshal(boardData)
	if err != nil {
		log.Printf("error serializing the response data")
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Write(response)
}
