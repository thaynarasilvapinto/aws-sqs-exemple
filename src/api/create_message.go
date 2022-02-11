package api

import (
	"src/config"
	"src/sqs"

	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateMessage(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		sqs.PublishMessage(string(body), config)

		var response map[string]interface{}
		json.Unmarshal([]byte(`{ "message": "mensagem publicada com sucesso" }`), &response)
		respondWithJSON(w, http.StatusOK, response)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
