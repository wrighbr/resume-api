package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wrighbr/resume-api/models"
)

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post
var contactInfo models.ContactInfo

func getContactInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contactInfo)
}

// func getPosts(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(posts)
// }
// func createPost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var post Post
// 	_ = json.NewDecoder(r.Body).Decode(&post)
// 	post.ID = strconv.Itoa(rand.Intn(1000000))
// 	posts = append(posts, post)
// 	json.NewEncoder(w).Encode(&post)
// }
// func getPost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for _, item := range posts {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Post{})
// }

func updateContantInfo(w http.ResponseWriter, r *http.Request) {
	var jsonData models.ContactInfo
	json.NewDecoder(r.Body).Decode(&jsonData)

	contactInfo = models.ContactInfo{
		Name:    jsonData.Name,
		Email:   jsonData.Email,
		Mobile:  jsonData.Mobile,
		Address: jsonData.Address,
		Town:    jsonData.Town,
		Country: jsonData.Country,
		Github:  jsonData.Github,
		Website: jsonData.Website,
	}

	fmt.Println("Print info here")
	fmt.Println(contactInfo)
	json.NewEncoder(w).Encode(contactInfo)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			var post Post
			_ = json.NewDecoder(r.Body).Decode(&post)
			post.ID = params["id"]
			posts = append(posts, post)
			json.NewEncoder(w).Encode(&post)
			return
		}
	}
	json.NewEncoder(w).Encode(posts)
}
func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
}
func main() {
	router := mux.NewRouter()
	// contactInfo = models.ContactInfo{
	// 	Mobile:  "123456789",
	// 	Name:    "John Doe",
	// 	Address: "123 fake st",
	// 	Town:    "springfield",
	// }
	posts = append(posts, Post{ID: "1", Title: "My first post", Body: "This is the content of my first post"})
	router.HandleFunc("/contactinfo", getContactInfo).Methods("GET")
	router.HandleFunc("/contactinfo", updateContantInfo).Methods("PUT")
	// router.HandleFunc("/posts", getPosts).Methods("GET")
	// router.HandleFunc("/posts", createPost).Methods("POST")
	// router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
