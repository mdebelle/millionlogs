package data

import (
	"encoding/json"
	"sync"
	"sync/atomic"
	"time"

	"github.com/mdebelle/milionlogs/query"
)

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

func New(modif time.Time) Data { return &data{seen: &sync.Map{}, filechange: modif} }

func (d *data) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Count int32 `json:"count"`
	}{d.count})
}

func (d *data) Count() int32                  { return d.count }
func (d *data) Filechange() time.Time         { return d.filechange }
func (d *data) SetLastFileChange(t time.Time) { d.filechange = t }
func (d *data) LoadOrStore(line interface{}, count *int32) {
	if v, ok := d.seen.LoadOrStore(line, count); ok {
		atomic.AddInt32(v.(*int32), 1)
	} else {
		atomic.AddInt32(&d.count, 1)
	}
}
