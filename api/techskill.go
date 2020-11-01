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

var techSkills []models.TechSkills
var colTechSkills = "techSkills"

func fillInTechSkills(r *http.Request) *models.TechSkills {
	var jsonData models.TechSkills
	json.NewDecoder(r.Body).Decode(&jsonData)

	return &models.TechSkills{
		ID:    rand.Intn(1000000),
		Skill: jsonData.Skill,
		Stars: jsonData.Stars,
		Tags:  jsonData.Tags,
	}
}

// createTechSkills godoc
// @Summary Creates Technical skill
// @Tags TechSkills
// @Accept  json
// @Produce  json
// @Success 200
// @Security BasicAuth
// @Router /techskill [post]
func createTechSkills(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	techSkills := fillInTechSkills(r)

	client.CreateDocument(colTechSkills, techSkills)

	location := "/techskill/" + strconv.Itoa(techSkills.ID)
	w.Header().Set("location", location)
}

// readTechSkills godoc
// @Summary Readss Technical skill
// @Tags TechSkills
// @Accept  json
// @Produce  json
// @Success 200
// @Router /techskill/{id} [get]
func readTechSkills(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params["ID"])
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		fmt.Println(err)
	}
	_, jsonData := client.ReadDocument(colTechSkills, id)

	json.NewEncoder(w).Encode(jsonData)
}

// updateTechSkills godoc
// @Summary Update Technical skill
// @Tags TechSkills
// @Accept  json
// @Produce  json
// @Success 200
// @Security BasicAuth
// @Router /techskill [put]
func updateTechSkills(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params["ID"])
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		fmt.Println(err)
	}

	jsonData := fillInTechSkills(r)
	docID, _ := client.ReadDocument(colTechSkills, id)

	client.UpdateDocument(colTechSkills, docID, jsonData)
}

// deleteTechSkills godoc
// @Summary Deletes Technical skill
// @Tags TechSkills
// @Accept  json
// @Produce  json
// @Success 200
// @Security BasicAuth
// @Router /techskill [delete]
func deleteTechSkills(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params["ID"])
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		fmt.Println(err)
	}

	docID, _ := client.ReadDocument(colTechSkills, id)

	client.DeleteDocument(colTechSkills, docID)
}

func getAllTechSkills(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	techSkills = nil
	var data models.TechSkills

	ctx := context.Background()
	client := createClient(ctx)

	iter := client.Collection(colTechSkills).Documents(ctx)
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
		techSkills = append(techSkills, data)
	}

	if jsonResponse {
		json.NewEncoder(w).Encode(techSkills)
	}
}
