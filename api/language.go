package api

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/wrighbr/resume-api/client"
	"github.com/wrighbr/resume-api/models"
	"google.golang.org/api/iterator"
)

var language []models.Language
var colLanguage = "language"

func fillInLanguage(r *http.Request) *models.Language {
	var jsonData models.Language
	json.NewDecoder(r.Body).Decode(&jsonData)

	return &models.Language{
		ID:          rand.Intn(1000000),
		Language:    jsonData.Language,
		Proficiency: jsonData.Proficiency,
	}
}

func createLanguage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	language := fillInLanguage(r)

	client.CreateDocument(colLanguage, language)

}

func readLanguage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params["ID"])
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		fmt.Println(err)
	}
	_, jsonData := client.ReadDocument(colLanguage, id)

	json.NewEncoder(w).Encode(jsonData)
}

func updateLanguage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params["ID"])
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		fmt.Println(err)
	}

	jsonData := fillInLanguage(r)
	docID, _ := client.ReadDocument(colLanguage, id)

	client.UpdateDocument(colLanguage, docID, jsonData)
}

func deleteLanguage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params["ID"])
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		fmt.Println(err)
	}

	docID, _ := client.ReadDocument(colLanguage, id)

	client.DeleteDocument(colLanguage, docID)
}

func getAllLanguage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	language = nil
	var data models.Language

	ctx := context.Background()
	client := createClient(ctx)

	iter := client.Collection(colLanguage).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		err = mapstructure.Decode(doc.Data(), &data)
		if err != nil {
			fmt.Println(err)
		}
		language = append(language, data)
	}

	if jsonResponse {
		json.NewEncoder(w).Encode(language)
	}
}
