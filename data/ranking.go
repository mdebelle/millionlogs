package data

import (
	"sort"
	"sync/atomic"

	"github.com/mdebelle/millionlogs/query"
)

func (d *data) Ranking() int {
	var insered = len(d.sorted)
	if insered == 0 {
		d.sorted = make([]*query.Query, d.count)
		d.seen.Range(func(key, value interface{}) bool {
			d.sorted[insered] = &query.Query{key.(string), atomic.LoadInt32(value.(*int32))}
			insered++
			return true
		})
		sort.Sort(query.Queries(d.sorted))
	}
	return insered
}

func (d *data) Popular(size int) []*query.Query {

	if size <= 0 {
		return []*query.Query{}
	}
	if insered := d.Ranking(); insered <= size {
		return d.sorted
	}
	return d.sorted[:size]
}
