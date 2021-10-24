package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"qrest/internal"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", internal.GetOne).Methods(http.MethodGet)
	router.HandleFunc("/rand", internal.GetRandomOne).Methods(http.MethodGet)
	router.HandleFunc("/custom", internal.GetCustomOne).Methods(http.MethodPost)
	router.HandleFunc("/doc", internal.Info).Methods(http.MethodGet)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})

	apiPort, ok := os.LookupEnv("PORT")
	if !ok {
		apiPort = "8088"
	}
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf("0.0.0.0:%s", apiPort),
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
