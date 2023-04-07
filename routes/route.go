package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DedeMarantes/times/controllers"
	"github.com/DedeMarantes/times/middleware"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	router := mux.NewRouter()
	router.Use(middleware.SetHeader)
	fmt.Println("Iniciando Servidor em http://localhost:8000")
	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/times", controllers.GetAllTimes).Methods("GET")
	router.HandleFunc("/jogadores", controllers.GetAllJogadores).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
