package store_test

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/mdebelle/millionlogs/store"
)

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
