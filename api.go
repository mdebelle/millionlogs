package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func writeJson(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}

func chronoRequest(r *http.Request, start time.Time) {
	fmt.Printf(
		"[%s] %s took %v\n",
		r.Method, r.URL.String(), time.Since(start),
	)
}

func popular(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer chronoRequest(r, start)

	if r.Method != http.MethodGet {
		writeJson(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	}

	date := strings.Replace(r.URL.Path, "/1/queries/popular/", "", 1)
	size := r.URL.Query().Get("size")
	v, status, err := loadOrCreate(date)
	if err != nil {
		writeJson(w, err.Error(), status)
		return
	}
	var max int
	if size == "" {
		max = 0
	} else if max, err = strconv.Atoi(size); err != nil {
		writeJson(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJson(w, v.Popular(max), status)
}

func count(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer chronoRequest(r, start)

	if r.Method != http.MethodGet {
		writeJson(w, http.StatusText(http.StatusNotImplemented), http.StatusNotImplemented)
		return
	}

	date := strings.Replace(r.URL.Path, "/1/queries/count/", "", 1)
	v, status, err := loadOrCreate(date)
	if err != nil {
		writeJson(w, err.Error(), status)
		return
	}
	writeJson(w, v, status)
}
