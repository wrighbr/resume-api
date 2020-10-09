package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {

	router := mux.NewRouter()

	router.HandleFunc("/contactinfo", getContactInfo).Methods("GET")
	router.HandleFunc("/contactinfo", updateContantInfo).Methods("PUT")
	router.HandleFunc("/contactinfo", deleteContantInfo).Methods("PUT")

	router.HandleFunc("/experience", createExperience).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}
