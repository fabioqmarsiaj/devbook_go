package controllers

import "net/http"

// insert an user into DB
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user..."))
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