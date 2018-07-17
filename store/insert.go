package store

import (
	"log"
	"regexp"
	"time"

	"github.com/mdebelle/millionlogs/data"
)

func insert(date, query string, modif time.Time) {
	v, ok := requests.Load(date)
	if !ok {
		v = data.New(modif)
		requests.Store(date, v)
	}
	var count int32 = 1
	v.(data.Data).LoadOrStore(query, &count)

}

type insertFunc func(*regexp.Regexp, time.Time, []byte)

// Insert type insertFunc create or update Data only for the requested date
func Insert(reg *regexp.Regexp, modif time.Time, logline []byte) {

	matches := reg.FindAllSubmatch(logline, -1)
	if len(matches) == 0 || len(matches[0]) != 3 {
		log.Printf("logline discarded regexp did not match correctly")
		return
	}
	// date, query, modif
	insert(string(matches[0][1]), string(matches[0][2]), modif)
}

// InsertAll type insertFunc create or update Data for all date possibility in the logline
func InsertAll(reg *regexp.Regexp, modif time.Time, logline []byte) {

	matches := reg.FindAllSubmatch(logline, -1)
	if len(matches) == 0 || len(matches[0]) != 8 {
		log.Printf("logline discarded regexp did not match correctly")
		return
	}
	length := len(matches[0])
	query := string(matches[0][length-1])
	for i := length - 2; i >= 0; i-- {
		// date, query, modif
		insert(string(matches[0][i]), query, modif)
	}
}
