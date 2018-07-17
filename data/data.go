package data

import (
	"encoding/json"
	"sync"
	"sync/atomic"
	"time"

	"github.com/mdebelle/millionlogs/query"
)

// Data interface to manipulate data
type Data interface {
	MarshalJSON() ([]byte, error)
	Count() int32
	Filechange() time.Time
	SetLastFileChange(time.Time)
	LoadOrStore(line interface{}, count *int32)

	Ranking() int
	Popular(size int) []*query.Query
}

type data struct {
	count      int32
	seen       *sync.Map
	filechange time.Time
	sorted     []*query.Query
}

// New return a new *data already setup which implement Data interface
func New(modif time.Time) Data { return &data{seen: &sync.Map{}, filechange: modif} }

// MarshalJSON used during marshaling a data
func (d *data) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Count int32 `json:"count"`
	}{d.count})
}

// Count return the number of query associated to the data
func (d *data) Count() int32 { return d.count }

// Filechange returns the base time of the file used to define the data
func (d *data) Filechange() time.Time { return d.filechange }

// SetLastFileChange record the base time used
func (d *data) SetLastFileChange(t time.Time) { d.filechange = t }

// LoadOrStore load or store a query counter
// if the query exist the query's counter is incremented
// if the query doesn't exist the data's counter is incremented
func (d *data) LoadOrStore(query interface{}, count *int32) {
	if v, ok := d.seen.LoadOrStore(query, count); ok {
		atomic.AddInt32(v.(*int32), 1)
	} else {
		atomic.AddInt32(&d.count, 1)
	}
}
