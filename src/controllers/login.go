package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"webapp/src/models"
	"webapp/src/responses"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.ResponseJON(w, http.StatusBadRequest, responses.ErrorAPI{Error: err.Error()})
		return
	}

	resp, err := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		responses.ResponseJON(w, http.StatusInternalServerError, responses.ErrorAPI{Error: err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		responses.VerifyStatusCodeErrors(w, resp)
		return
	}

	var data models.AuthDTO
	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
		responses.ResponseJON(w, http.StatusUnprocessableEntity, responses.ErrorAPI{Error: err.Error()})
		return
	}

	responses.ResponseJON(w, http.StatusOK, nil)
}
