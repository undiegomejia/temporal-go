package main

import (
	"fmt"
	"net/http"
)

func hairHandler(w http.ResponseWriter, r *http.Request) {
	_, eyes_ok := r.URL.Query()["eyes"]
	_, ears_ok := r.URL.Query()["ears"]
	_, mouth_ok := r.URL.Query()["mouth"]
	if eyes_ok && ears_ok && mouth_ok {
		hair := "black hair"
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, hair)
	} else {
		http.Error(w, "Missing required face parts.", http.StatusBadRequest)
	}
}

func voiceHandler(w http.ResponseWriter, r *http.Request) {
	_, nose_ok := r.URL.Query()["nose"]
	_, hair_ok := r.URL.Query()["hair"]
	if nose_ok && hair_ok {
		voice := "big voice"
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, voice)
	} else {
		http.Error(w, "Missing required face parts.", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/hair", hairHandler)
	http.HandleFunc("/voice", voiceHandler)
	http.ListenAndServe(":9999", nil)
}
