package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// insert an user into DB
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro	= json.Unmarshal(requestBody, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := db.Connect()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositories.NewUserRepositories(db)
	repository.Create(user)

}
// search all users on DB
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching all user..."))
}
// search one user on DB
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Searching an user..."))
}
// update an user on DB
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update an user..."))
}
// delete an user on DB
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting an user..."))
}