package store

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/mdebelle/milionlogs/data"
)

var requests = &sync.Map{}

func RankingAll() {
	requests.Range(func(key, value interface{}) bool {
		value.(data.Data).Ranking()
		return true
	})
}

type insertFunc func(*regexp.Regexp, time.Time, []byte)

func scan(reg *regexp.Regexp, modif time.Time, f insertFunc) error {
	content, err := os.Open(*Sample)
	if err != nil {
		log.Fatal(err)
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	scanner.Buffer([]byte{}, 4096)

	for scanner.Scan() {
		f(reg, modif, scanner.Bytes())
	}
	return scanner.Err()
}
