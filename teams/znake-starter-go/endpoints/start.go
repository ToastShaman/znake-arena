package endpoints

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Start(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var startRequest StartRequest
	err := json.NewDecoder(r.Body).Decode(&startRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	println(startRequest.GameID)

	var startResponse = StartResponse{"#123456", "#654321", "", "", "pixel", "pixel"}
	w.WriteHeader(200)

	response, err := json.MarshalIndent(startResponse, "", "  ")

	fmt.Fprintf(w, string(response))
}

type StartRequest struct {
	GameID int `json:"game_id"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type StartResponse struct {
	Color          string `json:"color"`
	SecondaryColor string `json:"secondary_color"`
	HeadURL        string `json:"head_url"`
	Taunt          string `json:"taunt"`
	HeadType       string `json:"head_type"`
	TailType       string `json:"tail_type"`
}