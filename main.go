package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/mdebelle/millionlog/data"
	"github.com/mdebelle/millionlog/store"
)

var port = flag.String("port", ":8080", "select your favorite port to expose the api")
var preload = flag.Bool("preload", false, "preload all data for faster requests but longer init")
var prerank = flag.Bool("prerank", false, "prerank all data for faster requests but longer init")
var sample = flag.String("sample", "sampledata/hn_logs.tsv", "set your sample (.tsv) file path location ")

func main() {
	flag.Parse()

	store.Prerank = prerank
	store.Sample = sample

	fmt.Println("server init...")
	start := time.Now()
	if *preload {
		info, err := os.Lstat(sample)
		if err != nil {
			panic(err)
		}
		if err := store.Scan(r, info.ModTime(), store.InsertAll); err != nil {
			panic(err)
		}
	}
	if *prerank {
		store.RankAllData()
	}
	fmt.Println("initialisation took", time.Since(start))
	fmt.Println("server ready on port", *port)

	http.HandleFunc("/1/queries/count/", count)
	http.HandleFunc("/1/queries/popular/", popular)
	log.Fatal(http.ListenAndServe(*port, nil))
}
