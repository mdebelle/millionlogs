package store

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/mdebelle/millionlogs/data"
)

var requests = &sync.Map{}

// Requests accesible outside of pkg
func Requests() *sync.Map { return requests }

// ResetRequests set to zero requests content
func ResetRequests() { requests = &sync.Map{} }

// RankingAll apply the Ranking method for each Data stored
func RankingAll() {
	requests.Range(func(key, value interface{}) bool {
		value.(data.Data).Ranking()
		return true
	})
}

// Scan the selected file which contains tone of logs
func Scan(reg *regexp.Regexp, modif time.Time, f InsertFunc) error {
	if Sample == nil {
		return fmt.Errorf("could not open: file sample not set")
	}

	content, err := os.Open(*Sample)
	if err != nil {
		return err
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	scanner.Buffer([]byte{}, 4096)

	for scanner.Scan() {
		f(reg, modif, scanner.Bytes())
	}
	return scanner.Err()
}
