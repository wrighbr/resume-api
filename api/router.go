package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {

	router := mux.NewRouter()

	router.HandleFunc("/contactinfo", getContactInfo).Methods("GET")
	router.HandleFunc("/contactinfo", createContantInfo).Methods("POST")
	router.HandleFunc("/contactinfo", updateContantInfo).Methods("PUT")
	router.HandleFunc("/contactinfo", deleteContantInfo).Methods("DELETE")

	router.HandleFunc("/experience", getAllExperience).Methods("GET")
	router.HandleFunc("/experience", createExperience).Methods("POST")
	router.HandleFunc("/experience/{ID}", readExperience).Methods("GET")
	router.HandleFunc("/experience/{ID}", updateExperience).Methods("PUT")
	router.HandleFunc("/experience/{ID}", deleteExperience).Methods("DELETE")

	router.HandleFunc("/education", getAllEducation).Methods("GET")
	router.HandleFunc("/education", createEducation).Methods("POST")
	router.HandleFunc("/education/{ID}", readEducation).Methods("GET")
	router.HandleFunc("/education/{ID}", updateEducation).Methods("PUT")
	router.HandleFunc("/education/{ID}", deleteEducation).Methods("DELETE")

	router.HandleFunc("/resume", getResume).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
