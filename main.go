package main

import (
	"github.com/gorilla/mux"
	"github.com/mainak90/simple-urlshortner/driver"
	"github.com/mainak90/simple-urlshortner/handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	db, err := driver.InitDB()
	if err != nil {
		panic(err)
	}
	dbclient := handlers.DBClient{Db: db}
	defer db.Close()
	r := mux.NewRouter()
	r.HandleFunc("/heartbeat", handlers.Heartbeat).Methods("GET")
	r.HandleFunc("/v1/short/{encoded_string:[a-zA-Z0-9]*}", dbclient.GenerateShortURL).Methods("GET")
	r.HandleFunc("/v1/short", dbclient.GenerateShortURL).Methods("POST")
	server := &http.Server{
		Addr:              "127.0.0.1:8011",
		Handler:           r,
		TLSConfig:         nil,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
