package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "login.html", nil)
}

func SignupPage(w http.ResponseWriter, r *http.Request) {
	utils.Render(w, "signup.html", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.APIURL)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodGet, url, nil)
	if err != nil {
		responses.ResponseJON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		responses.VerifyStatusCodeErrors(w, resp)
		return
	}

	var posts []models.PostDTO
	if err = json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		responses.ResponseJON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	utils.Render(w, "home.html", posts)
}
