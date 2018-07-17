package query

// Query contains
// Query as the log line
// Count as the number of times Query appears in the logsfile
type Query struct {
	Query string `json:"query"`
	Count int32  `json:"count"`
}

// Queries is an alias for a slice of pointer to query
// Queries implement sort interface
type Queries []*Query

func (q Queries) Len() int           { return len(q) }
func (q Queries) Less(i, j int) bool { return q[i].Count > q[j].Count }
func (q Queries) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
