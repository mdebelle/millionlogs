package store

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mdebelle/millionlogs/data"
)

// Sample is the flag which contain the file path
var Sample *string

// Prerank is the flag option
var Prerank *bool

// LoadOrCreate try to load Data associated to the date parameter
// If no Data found it will try to create it
// this function take care of the file changement
func LoadOrCreate(date string) (data.Data, int, error) {

	info, err := os.Lstat(*Sample)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	filechange := info.ModTime()

	if d, ok := requests.Load(date); ok {
		if filechange == d.(data.Data).Filechange() {
			return d.(data.Data), http.StatusOK, nil
		}
		requests.Delete(date)
	}
	reg, status, err := buildRegexp(date)
	if err != nil {
		return nil, status, err
	}
	if err := Scan(reg, filechange, Insert); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	if d, ok := requests.Load(date); ok {
		if *Prerank {
			d.(data.Data).Ranking()
		}
		return d.(data.Data), http.StatusOK, nil
	}
	return nil, http.StatusNotFound, fmt.Errorf("no data found for '%s'", date)
}
