package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloworldResponse struct {
	Message string `json:"message"`
	// do not output this field
	// Author string `json:"-"`
	// do not output the field if the value empty
	// Date string `json:",omitempty"`
	// convert out to a string and rename "id"
	// ID int `json:"id,string"`
}

type helloworldRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 8080

	cathandler := http.FileServer(http.Dir("./images"))
	http.Handle("/cat/", http.StripPrefix("/cat/", cathandler))

	http.HandleFunc("/helloworld", helloworldHandler)

	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloworldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloworldRequest
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	response := helloworldResponse{Message: "Hello " + request.Name}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}
