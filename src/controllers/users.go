package controllers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ResponseJON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users", config.APIURL)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		responses.ResponseJON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		responses.VerifyStatusCodeErrors(w, resp)
		return
	}

	responses.ResponseJON(w, resp.StatusCode, nil)

}

func UnfollowerUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.ResponseJON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/unfollower", config.APIURL, userID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodPost, url, nil)
	if err != nil {
		responses.ResponseJON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		responses.VerifyStatusCodeErrors(w, resp)
		return
	}

	responses.ResponseJON(w, resp.StatusCode, nil)
}

func FollowerUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.ResponseJON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/users/%d/follower", config.APIURL, userID)
	resp, err := requests.MakeRequestWithAuthentication(r, http.MethodPost, url, nil)
	if err != nil {
		responses.ResponseJON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		responses.VerifyStatusCodeErrors(w, resp)
		return
	}

	responses.ResponseJON(w, resp.StatusCode, nil)
}
