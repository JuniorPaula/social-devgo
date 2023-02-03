package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "login.html", nil)
}

func SignupPage(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "signup.html", nil)
}
