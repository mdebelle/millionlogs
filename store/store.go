package store

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mdebelle/milionlogs/data"
)

var Sample *string
var Prerank *bool

func loadOrCreate(date string) (data.Data, int, error) {

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
	if err := scan(reg, filechange, insert); err != nil {
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
