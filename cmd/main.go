package main

import (
	"log"
	"net/http"

	"github.com/Yefhem/basic-api-with-go/pkg/db"
	"github.com/Yefhem/basic-api-with-go/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {

	DB := db.Init()
	h := handlers.New(DB)

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/albums", h.GetAllAlbumHanlder).Methods("GET")
	router.HandleFunc("/api/v1/albums/{id:[0-9]+}", h.GetAlbumHandler).Methods("GET")

	router.HandleFunc("/api/v1/albums", h.AddAlbumHandler).Methods("POST")
	router.HandleFunc("/api/v1/albums/{id:[0-9]+}", h.PutAlbumHandler).Methods("PUT")
	router.HandleFunc("/api/v1/albums/{id:[0-9]+}", h.DeleteAlbumHandler).Methods("DELETE")

	log.Println("API is running!")
	http.ListenAndServe(":9000", router)
}
