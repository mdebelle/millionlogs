package store_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/mdebelle/millionlogs/data"
	"github.com/mdebelle/millionlogs/store"
)

var prerankTrue = true

func donothing() {}
func prerank()   { store.Prerank = &prerankTrue }
func justReset() { store.ResetRequests() }
func setuprequests() {
	store.ResetRequests()
	store.InsertAll(store.RegAll, date, []byte(`2015-08-01 00:03:43	http%3A%2F%2Ftechacute.com%2F10-essentials-every-desk-needs%2F`))
}

// func TestLoadOrCreate(date string) (data.Data, int, error)
func TestLoadOrCreate(t *testing.T) {
	sampleName := "../sample/small.tsv"
	sampleFake := "fakenews"
	type expected struct {
		Data   data.Data
		Status int
		ErrStr string
	}
	testTable := []struct {
		Name     string
		Date     string
		sample   *string
		f        func() // preset requests
		Expected expected
	}{
		{"sample nil", "2015", nil, donothing, expected{nil, http.StatusInternalServerError, "could not open: file sample not set"}},
		{"sample doesn't exist", "2015", &sampleFake, donothing, expected{nil, http.StatusInternalServerError, "lstat fakenews: no such file or directory"}},
		{"regexp fail", "hello world", &sampleName, setuprequests, expected{nil, http.StatusBadRequest, "could not use 'hello world' as a date parameter: bad format"}},
		{"prerank", "2015", &sampleName, prerank, expected{nil, http.StatusOK, "<nil>"}},
		{"not found", "2016", &sampleName, justReset, expected{nil, http.StatusNotFound, "no data found for '2016'"}},
	}

	store.ResetRequests()
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			testCase.f()
			store.Sample = testCase.sample
			_, s, err := store.LoadOrCreate(testCase.Date)
			if s != testCase.Expected.Status {
				t.Fatalf("expected status %d got %d", testCase.Expected.Status, s)
			}
			if errStr := fmt.Sprintf("%v", err); errStr != testCase.Expected.ErrStr {
				t.Fatalf("expected err %s got %s", testCase.Expected.ErrStr, errStr)
			}
		})
	}
}
