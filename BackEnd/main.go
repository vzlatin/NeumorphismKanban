package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vzlatin/NeumorphismKanban/httphandlers"
	ws "github.com/vzlatin/NeumorphismKanban/wsconnection"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	PORT := os.Getenv("PORT")

	setupApi()
	fmt.Println("Server started on port: ", PORT)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupApi() {

	manager := ws.NewManager()
	http.HandleFunc("/ws", manager.ServeWS)
	http.HandleFunc("/getBoardData", httphandlers.GetBoardIdData)

	// will probably be removed
	http.HandleFunc("/getBoards", httphandlers.GetAllBoards) 
}
