package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Numbers struct {
	A int `json:"a"`
	B int `json:"b"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/addition", add()).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/subtract", sub).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/multiply/{id}/{di}", mul).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func add() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var reqBdy Numbers

		err := json.NewDecoder(req.Body).Decode(&reqBdy)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		sum := reqBdy.A + reqBdy.B
		response, err := json.Marshal(sum)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func sub(w http.ResponseWriter, r *http.Request) {
	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")

	first, err := strconv.Atoi(a)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	second, err := strconv.Atoi(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	subtraction := first - second
	response, err := json.Marshal(subtraction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func mul(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	di := vars["di"]
	fmt.Println(vars)
	w.Write([]byte(id + " " + di))
}
