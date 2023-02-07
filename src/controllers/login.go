package controllers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
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
	token, _ := io.ReadAll(resp.Body)
	fmt.Println("Resposta ", resp.StatusCode, string(token))
}
