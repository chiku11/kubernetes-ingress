package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

type Location struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func health(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
}

func findLocation(response http.ResponseWriter, request *http.Request) {
	var p Person
	var l Location
	personState := make(map[string]string)

	personState["srikant"] = "elgin IL"
	personState["siri"] = "schaumburg IL"
	err := json.NewDecoder(request.Body).Decode(&p)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	split := strings.Split(personState[p.Name], " ")
	l.City = split[0]
	l.State = split[1]
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(l)
}

func main() {
	http.HandleFunc("/", health)
	http.HandleFunc("/location", findLocation)
	http.ListenAndServe(":8091", nil)
}
