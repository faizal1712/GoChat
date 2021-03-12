package routes

import (
	"contact/socket"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func CreateRouter() {
	router = mux.NewRouter()

}

func InititalizeRoutes() {
	router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./asset"))))
	router.Handle("/socket.io/", socket.Server)
}

func StartServer() {
	fmt.Println("Server Started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
