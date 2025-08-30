package controllers

import {
	"net/http"
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid Request Method", er)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, ok := users[username] // add in memory database for users

	if !ok || !checkPasswordHash(password, user.PassWord){
		er := http.StatusUnauthorized
		http.Error(w, "Invalid username or password", er)
		return
	}

	// generate session token

	sessionToken := generateSessionToken(32)

	http.SetCookie(w, &http.Cookie)

	fmt.Fprintln(w, "Login successful!")

}
func GetUser(w http.ResponseWriter, r *http.Request)    {}
func GetUsers(w http.ResponseWriter, r *http.Request)   {}
func CreateUser(w http.ResponseWriter, r *http.Request) {}
func UpdateUser(w http.ResponseWriter, r *http.Request) {}
func DeleteUser(w http.ResponseWriter, r *http.Request) {}
