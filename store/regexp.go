package store

import (
	"fmt"
	"net/http"
	"regexp"
)

const (
	year   = 4
	month  = 7
	day    = 10
	hour   = 13
	minute = 16
	second = 19
)

// RegAll is a regular expression that contain a group for each possible date
// composed of 6 groups for date and 1 group for the query
var RegAll = regexp.MustCompile(`(?m)^((((((\d\d\d\d)-\d\d)-\d\d) \d\d):\d\d):\d\d)\t(.*)$`)

var acceptedDateFMT = regexp.MustCompile(
	`(?m)^(((\d{4})(-\d\d){2} \d\d(:\d\d){2})` +
		`|((\d{4})(-\d\d){2} \d\d(:\d\d))` +
		`|((\d{4})(-\d\d){2} \d\d)` +
		`|((\d{4})(-\d\d){2})` +
		`|((\d{4})(-\d\d))` +
		`|((\d{4})))$`,
)

func BuildRegexp(date string) (*regexp.Regexp, int, error) {

	if !acceptedDateFMT.MatchString(date) {
		return nil, http.StatusBadRequest, fmt.Errorf("could not use '%s' as a date parameter: bad format", date)
	}

	switch len(date) {
	case year:
		return regexp.MustCompile(`(?m)^(` + date + `)-\d\d-\d\d \d\d:\d\d:\d\d\t(.*)$`), http.StatusOK, nil
	case month:
		return regexp.MustCompile(`(?m)^(` + date + `)-\d\d \d\d:\d\d:\d\d\t(.*)$`), http.StatusOK, nil
	case day:
		return regexp.MustCompile(`(?m)^(` + date + `) \d\d:\d\d:\d\d\t(.*)$`), http.StatusOK, nil
	case hour:
		return regexp.MustCompile(`(?m)^(` + date + `):\d\d:\d\d\t(.*)$`), http.StatusOK, nil
	case minute:
		return regexp.MustCompile(`(?m)^(` + date + `):\d\d\t(.*)$`), http.StatusOK, nil
	default:
		return regexp.MustCompile(`(?m)^(` + date + `)\t(.*)$`), http.StatusOK, nil
	}
}
