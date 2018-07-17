package data_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mdebelle/millionlogs/data"
)

var date = time.Date(2018, 7, 17, 22, 32, 17, 0, time.UTC)
var dateChange = time.Date(2018, 7, 17, 23, 42, 0, 0, time.UTC)

func NewData(lines []string) data.Data {
	d := data.New(date)
	for _, v := range lines {
		var count int32 = 1
		d.LoadOrStore(v, &count)
	}
	return d
}

func TestMarshalJSON(t *testing.T) {
	type expected struct {
		str string
		err string
	}
	testTable := []struct {
		Name     string
		D        data.Data
		Expected expected
	}{
		{"new", data.New(date), expected{`{"count":0}`, "<nil>"}},
		{"contain three queries (different)", NewData([]string{"un", "deux", "trois"}), expected{`{"count":3}`, "<nil>"}},
		{"contain three queries (identical)", NewData([]string{"un", "un", "un"}), expected{`{"count":1}`, "<nil>"}},
	}
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			b, err := testCase.D.MarshalJSON()
			if string(b) != testCase.Expected.str {
				t.Fatalf("expected content %s got %s", testCase.Expected.str, string(b))
			}
			if fmt.Sprintf("%v", err) != testCase.Expected.err {
				t.Fatalf("expected content %s got %v", testCase.Expected.err, err)
			}

		})
	}
}

func TestCount(t *testing.T) {

	testTable := []struct {
		Name     string
		D        data.Data
		Expected int32
	}{
		{"new", data.New(date), 0},
		{"contain three queries (different)", NewData([]string{"un", "deux", "trois"}), 3},
		{"contain three queries (identical)", NewData([]string{"un", "un", "un"}), 1},
	}
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			c := testCase.D.Count()
			if c != testCase.Expected {
				t.Fatalf("expected content %d got %d", testCase.Expected, c)
			}
		})
	}
}
func TestFilechange(t *testing.T) {

	testTable := []struct {
		Name     string
		D        data.Data
		Expected time.Time
	}{
		{"new", data.New(date), date},
	}
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			c := testCase.D.Filechange()
			if c != testCase.Expected {
				t.Fatalf("expected content %v got %v", testCase.Expected, c)
			}
		})
	}
}
func TestSetLastFileChange(t *testing.T) {

	testTable := []struct {
		Name     string
		D        data.Data
		newDate  time.Time
		Expected time.Time
	}{
		{"same date", data.New(date), date, date},
		{"new date", data.New(date), dateChange, dateChange},
	}
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			testCase.D.SetLastFileChange(testCase.newDate)
			c := testCase.D.Filechange()
			if c != testCase.Expected {
				t.Fatalf("expected content %v got %v", testCase.Expected, c)
			}
		})
	}
}
