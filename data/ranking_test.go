package data_test

import (
	"encoding/json"
	"testing"

	"github.com/mdebelle/millionlogs/data"
)

func TestRanking(t *testing.T) {
	testTable := []struct {
		Name     string
		D        data.Data
		Expected int
	}{
		{"new", data.New(date), 0},
		{"contain three queries (different)", NewData([]string{"un", "deux", "trois"}), 3},
		{"contain three queries (identical)", NewData([]string{"un", "un", "un"}), 1},
	}
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			insered := testCase.D.Ranking()
			if insered != testCase.Expected {
				t.Fatalf("expected insered %d got %d", testCase.Expected, insered)
			}
		})
	}
}
func TestPopular(t *testing.T) {
	type expected struct {
		length int
		order  string
	}
	testTable := []struct {
		Name     string
		D        data.Data
		size     int
		Expected expected
	}{
		{"new", data.New(date), 2, expected{0, "[]"}},
		{"negative", data.New(date), -3, expected{0, "[]"}},
		{"contain three queries (identical)", NewData([]string{"un", "un", "un"}), 1,
			expected{1, `[{"query":"un","count":3}]`}},
		{"contain three queries (order)", NewData([]string{"un", "un", "deux"}), 2,
			expected{2, `[{"query":"un","count":2},{"query":"deux","count":1}]`}},
		{"contain more queries than size", NewData([]string{"un", "un", "un", "deux", "deux", "trois"}), 2,
			expected{2, `[{"query":"un","count":3},{"query":"deux","count":2}]`}},
		{"contain less queries than size", NewData([]string{"un", "un", "deux"}), 4,
			expected{2, `[{"query":"un","count":2},{"query":"deux","count":1}]`}},
	}
	for _, testCase := range testTable {
		t.Run(testCase.Name, func(t *testing.T) {
			queries := testCase.D.Popular(testCase.size)
			str, err := json.Marshal(queries)
			if err != nil {
				t.Errorf("fail to marshal response: %v", err)
			}
			if len(queries) != testCase.Expected.length {
				t.Fatalf("expected insered %d got %d", testCase.Expected.length, len(queries))
			}
			if string(str) != testCase.Expected.order {
				t.Fatalf("expected order %s got %s", testCase.Expected.order, string(str))
			}
		})
	}
}
