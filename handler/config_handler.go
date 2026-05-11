package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API radi")
}

type Config struct {
	Name    string            `json:"name"`
	Version string            `json:"version"`
	Params  map[string]string `json:"params"`
}

func CreateConfig(w http.ResponseWriter, r *http.Request) {

	var c Config
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "C primljen")
}
