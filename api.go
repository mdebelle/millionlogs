package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/mdebelle/millionlogs/store"
)

func writeJson(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

func chronoRequest(r *http.Request, start time.Time) {
	log.Printf(
		"[%s] %s took %v\n",
		r.Method, r.URL.String(), time.Since(start),
	)
}

func handler() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/1/queries/count/", count)
	r.HandleFunc("/1/queries/popular/", popular)
	return r
}

func popular(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer chronoRequest(r, start)

	if r.Method != http.MethodGet {
		writeJson(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	}

	date := strings.Replace(r.URL.Path, "/1/queries/popular/", "", 1)
	v, status, err := store.LoadOrCreate(date)
	if err != nil {
		writeJson(w, err.Error(), status)
		return
	}
	text := r.FormValue("size")
	if text == "" {
		http.Error(w, "missing size", http.StatusBadRequest)
		return
	}
	size, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "not a number: "+text, http.StatusBadRequest)
		return
	}

	writeJson(w, struct {
		Queries interface{} `json:"queries"`
	}{v.Popular(size)}, status)
}

func count(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer chronoRequest(r, start)

	if r.Method != http.MethodGet {
		writeJson(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	}

	date := strings.Replace(r.URL.Path, "/1/queries/count/", "", 1)
	v, status, err := store.LoadOrCreate(date)
	if err != nil {
		writeJson(w, err.Error(), status)
		return
	}
	writeJson(w, v, status)
}
