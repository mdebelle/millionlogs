package store

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/mdebelle/millionlogs/data"
)

var requests = &sync.Map{}

// RankingAll apply the Ranking method for each Data stored
func RankingAll() {
	requests.Range(func(key, value interface{}) bool {
		value.(data.Data).Ranking()
		return true
	})
}

type insertFunc func(*regexp.Regexp, time.Time, []byte)

// Scan the selected file which contains tone of logs
func Scan(reg *regexp.Regexp, modif time.Time, f insertFunc) error {
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
