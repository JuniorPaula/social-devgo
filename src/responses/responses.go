package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorAPI struct {
	Error string `json:"error"`
}

func ResponseJON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}

}

func VerifyStatusCodeErrors(w http.ResponseWriter, r *http.Response) {
	var err ErrorAPI
	json.NewDecoder(r.Body).Decode(&err)
	ResponseJON(w, r.StatusCode, err)
}
