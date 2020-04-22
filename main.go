package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/osmanzz/workshop-backend/service"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/select_all_data", service.GetAllData).Methods("GET")
	router.HandleFunc("/insert_data_mahasiswa", service.InsertMahasiswa).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	handler := c.Handler(router)

	log.Println("======== successfull listening to :8080")
	http.ListenAndServe(":8080", handler)
}
