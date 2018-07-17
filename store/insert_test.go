package store_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/mdebelle/millionlogs/data"
	"github.com/mdebelle/millionlogs/store"
)

var date = time.Date(2018, 7, 17, 22, 32, 17, 0, time.UTC)
var dateChange = time.Date(2018, 7, 17, 23, 42, 0, 0, time.UTC)

func TestInsert(t *testing.T) {
	type expected struct {
		queriesCount   int
		queriesCounter []int32
	}
	rg2015, _, _ := store.BuildRegexp("2015")
	testTable := []struct {
		Name     string
		reg      *regexp.Regexp
		modif    time.Time
		logline  []byte
		Expected expected
	}{
		{
			"new match",
			rg2015, date, []byte(`2015-08-01 00:03:43	http%3A%2F%2Ftechacute.com%2F10-essentials-every-desk-needs%2F`),
			expected{1, []int32{1}},
		}, {
			"match already exist",
			rg2015, date, []byte(`2015-08-01 00:03:43	http%3A%2F%2Ftechacute.com%2F10-essentials-every-desk-needs%2F`),
			expected{1, []int32{1}},
		}, {
			"match already exist time modified",
			rg2015, dateChange, []byte(`2015-08-01 00:03:43	http%3A%2F%2Ftechacute.com%2F10-essentials-every-desk-needs%2F`),
			expected{1, []int32{1}},
		}, {
			"regexp do not match correctly",
			store.RegAll, dateChange, []byte(`2015-08-01 00:03:43	http%3A%2F%2Ftechacute.com%2F10-essentials-every-desk-needs%2F`),
			expected{1, []int32{1}},
		},
	}

	store.ResetRequests()
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			store.Insert(testCase.reg, testCase.modif, testCase.logline)
			var i = 0
			store.Requests().Range(func(key, value interface{}) bool {
				if value.(data.Data).Count() != testCase.Expected.queriesCounter[i] {
					t.Errorf("queriescounter didn't match %d got %d", testCase.Expected.queriesCounter[i], value.(data.Data).Count())
					return false
				}
				i++
				return true
			})

			if i != testCase.Expected.queriesCount {
				t.Fatalf("queriescount didn't match %d got %d", testCase.Expected.queriesCount, i)
			}
		})
	}

}

func TestInsertAll(t *testing.T) {
	type expected struct {
		queriesCount   int
		queriesCounter []int32
	}
	rg2015, _, _ := store.BuildRegexp("2015")
	testTable := []struct {
		Name     string
		reg      *regexp.Regexp
		modif    time.Time
		logline  []byte
		Expected expected
	}{
		{
			"new match",
			store.RegAll, date, []byte(`2015-08-01 00:03:43	http%3A%2F%2Ftechacute.com%2F10-essentials-every-desk-needs%2F`),
			expected{7, []int32{1, 1, 1, 1, 1, 1, 1}},
		}, {
			"match already exist",
			store.RegAll, date, []byte(`2015-08-01 00:03:43	http%3A%2F%2Ftechacute.com%2F10-essentials-every-desk-needs%2F`),
			expected{7, []int32{1, 1, 1, 1, 1, 1, 1}},
		}, {
			"match already exist time modified",
			store.RegAll, dateChange, []byte(`2015-08-01 00:03:43	http%3A%2F%2Ftechacute.com%2F10-essentials-every-desk-needs%2F`),
			expected{7, []int32{1, 1, 1, 1, 1, 1, 1}},
		}, {
			"regexp do not match correctly",
			rg2015, dateChange, []byte(`2015-08-01 00:03:43	http%3A%2F%2Ftechacute.com%2F10-essentials-every-desk-needs%2F`),
			expected{7, []int32{1, 1, 1, 1, 1, 1, 1}},
		},
	}

	store.ResetRequests()
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			store.InsertAll(testCase.reg, testCase.modif, testCase.logline)
			var i = 0
			store.Requests().Range(func(key, value interface{}) bool {
				if value.(data.Data).Count() != testCase.Expected.queriesCounter[i] {
					t.Fatalf("queriescounter didn't match %d got %d", testCase.Expected.queriesCounter[i], value.(data.Data).Count())
					return false
				}
				i++
				return true
			})

			if i != testCase.Expected.queriesCount {
				t.Fatalf("queriescount didn't match %d got %d", testCase.Expected.queriesCount, i)
			}
		})
	}
}
