package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type rjson struct {
	Language string
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
	w.WriteHeader(http.StatusOK)
	return
}

func api(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Invalid request method.", 405)
		return
	}

	rj := &rjson{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&rj)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	data := ""
	if rj.Language == "python" {
		data = "That's nice, but slow!"
	} else {
		if rj.Language == "java" {
			data = "Faster than python!"
		} else {
			data = "I dont know!"
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
	return
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/api", api)
	print("I am online")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
