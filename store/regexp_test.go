package store_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/mdebelle/millionlogs/store"
)

func TestBuildRegexp(t *testing.T) {
	type expected struct {
		regStr string
		status int
		err    string
	}
	testTable := []struct {
		Name     string
		Date     string
		Expected expected
	}{
		{
			"not regular", "hello world",
			expected{"<nil>", http.StatusBadRequest, "could not use 'hello world' as a date parameter: bad format"},
		}, {
			"year regular", "2015",
			expected{`(?m)^(2015)-\d\d-\d\d \d\d:\d\d:\d\d\t(.*)$`, http.StatusOK, "<nil>"},
		}, {
			"month regular", "2015-08",
			expected{`(?m)^(2015-08)-\d\d \d\d:\d\d:\d\d\t(.*)$`, http.StatusOK, "<nil>"},
		}, {
			"day regular", "2015-08-01",
			expected{`(?m)^(2015-08-01) \d\d:\d\d:\d\d\t(.*)$`, http.StatusOK, "<nil>"},
		}, {
			"hour regular", "2015-08-01 00",
			expected{`(?m)^(2015-08-01 00):\d\d:\d\d\t(.*)$`, http.StatusOK, "<nil>"},
		}, {
			"minute regular", "2015-08-01 00:03",
			expected{`(?m)^(2015-08-01 00:03):\d\d\t(.*)$`, http.StatusOK, "<nil>"},
		}, {
			"second regular", "2015-08-01 00:03:43",
			expected{`(?m)^(2015-08-01 00:03:43)\t(.*)$`, http.StatusOK, "<nil>"},
		},
	}

	store.ResetRequests()
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			reg, status, err := store.BuildRegexp(testCase.Date)
			var regStr string
			if reg == nil {
				regStr = "<nil>"
			} else {
				regStr = reg.String()
			}

			if regStr != testCase.Expected.regStr {
				t.Fatalf("expected expression %s got %s", testCase.Expected.regStr, regStr)
			}
			if status != testCase.Expected.status {
				t.Fatalf("expected status %d got %d", testCase.Expected.status, status)
			}
			if errStr := fmt.Sprintf("%v", err); errStr != testCase.Expected.err {
				t.Fatalf("expected error %s got %s", testCase.Expected.err, errStr)
			}
		})
	}
}
