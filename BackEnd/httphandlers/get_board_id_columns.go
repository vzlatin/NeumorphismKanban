package httphandlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	ws "github.com/vzlatin/NeumorphismKanban/wsconnection"
)


func GetBoardIdColumns(w http.ResponseWriter, r *http.Request) {
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
	columns, err := apiConfig.DB.GetBoardIdColums(ctx, uuid)
	if err != nil {
		log.Printf("error: couldn't get the board id columns: %s", err)
	}

	response, err := json.Marshal(columns)
	if err != nil {
		log.Printf("error serializing the response data")
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Write(response)
}
