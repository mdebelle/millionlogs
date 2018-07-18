package store_test

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/mdebelle/millionlogs/store"
)

// []byte(`2015-08-01 00:03:43	http%3A%2F%2Ftechacute.com%2F10-essentials-every-desk-needs%2F`)

// Scan(reg *regexp.Regexp, modif time.Time, f insertFunc)
func TestScan(t *testing.T) {
	sampleName := "../sample/small.tsv"
	sampleFake := "fakenews"
	testTable := []struct {
		Name     string
		reg      *regexp.Regexp
		modif    time.Time
		f        store.InsertFunc
		sample   *string
		Expected string
	}{
		{"file exist", store.RegAll, date, store.InsertAll, &sampleName, "<nil>"},
		{"file doesn't exist", store.RegAll, date, store.InsertAll, &sampleFake, "open fakenews: no such file or directory"},
		{"file not set", store.RegAll, date, store.InsertAll, nil, "could not open: file sample not set"},
	}

	store.ResetRequests()
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			store.Sample = testCase.sample
			err := store.Scan(testCase.reg, testCase.modif, testCase.f)

			if errStr := fmt.Sprintf("%v", err); errStr != testCase.Expected {
				t.Fatalf("expected error %s got %s", testCase.Expected, errStr)
			}

		})
	}

}
