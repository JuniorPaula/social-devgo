package controllers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"webapp/src/config"
	"webapp/src/responses"
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
