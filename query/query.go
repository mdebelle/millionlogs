package query

type Query struct {
	Query string `json:"query"`
	Count int32  `json:"count"`
}

type Queries []*Query

func (q Queries) Len() int           { return len(q) }
func (q Queries) Less(i, j int) bool { return q[i].Count > q[j].Count }
func (q Queries) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
