package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		message := []byte("Sup!")
		w.Write(message)
	})
	http.HandleFunc("/ws", manager.ServeWS)

}
