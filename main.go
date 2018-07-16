package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mdebelle/millionlogs/store"
)

var port = flag.String("port", ":8080", "select your favorite port to expose the api")
var preload = flag.Bool("preload", false, "preload all data for faster requests but longer init")
var prerank = flag.Bool("prerank", false, "prerank all data for faster requests but longer init")
var sample = flag.String("sample", "sample/small.tsv", "set your sample (.tsv) file path location")

func main() {
	flag.Parse()

	store.Prerank = prerank
	store.Sample = sample

	log.Println("server init...")
	log.Println("selected file", *sample)
	start := time.Now()
	if *preload {
		info, err := os.Lstat(*sample)
		if err != nil {
			panic(err)
		}
		if err := store.Scan(store.RegAll, info.ModTime(), store.InsertAll); err != nil {
			panic(err)
		}
	}
	if *prerank {
		store.RankingAll()
	}
	log.Println("initialisation took", time.Since(start))
	log.Println("server ready on port", *port)

	http.HandleFunc("/1/queries/count/", count)
	http.HandleFunc("/1/queries/popular/", popular)
	log.Fatal(http.ListenAndServe(*port, nil))
}
