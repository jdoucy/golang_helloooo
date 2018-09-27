package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting test server for G")
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type request struct {
	Language string `json:"language"`
}

type response struct {
	Ok bool `json:"ok"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var request request
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, `{"error_message": "internal error"}`, http.StatusBadRequest)
		return
	}
	log.Printf("host: %v, referer: %v", r.Host, r.Referer())
	log.Printf("language: %v", request.Language)
	log.Printf("headers:")
	for i := range r.Header {
		log.Printf("\t%v => %v", i, r.Header.Get(i))
	}
	payload, err := json.Marshal(response{
		Ok: true,
	})
	if err != nil {
		http.Error(w, `{"error_message": "internal error"}`, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}
