package store

import (
	"regexp"
	"time"

	"github.com/mdebelle/milionlogs/data"
)

func insert(reg *regexp.Regexp, modif time.Time, logline []byte) {

	matches := reg.FindAllSubmatch(logline, -1)
	if len(matches) == 0 {
		return
	}
	date := string(matches[0][1])
	line := string(matches[0][2])
	v, ok := requests.Load(date)
	if !ok {
		v = data.New(modif)
		requests.Store(date, v)
	}
	var count int32 = 1
	v.(data.Data).LoadOrStore(line, &count)
}

func insertAll(reg *regexp.Regexp, modif time.Time, logline []byte) {

	matches := reg.FindAllSubmatch(logline, -1)
	if len(matches) == 0 {
		return
	}
	length := len(matches[0])
	line := string(matches[0][length-1])
	for i := length - 2; i >= 0; i-- {
		date := string(matches[0][i])
		v, ok := requests.Load(date)
		if !ok {
			v = data.New(modif)
			requests.Store(date, v)
		}
		var count int32 = 1
		v.(data.Data).LoadOrStore(line, &count)
	}
}
