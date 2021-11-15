package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

// Build returns a simple tree from a give array
// of records
func Build(records []Record) (*Node, error) {
	node := make(map[int]*Node, len(records))
	// Sort by ID.
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, errors.New("not in sequence or has bad parent")
		}
		node[r.ID] = &Node{ID: r.ID}
		if r.ID != 0 {
			node[r.Parent].Children = append(node[r.Parent].Children, node[r.ID])
		}
	}
	return node[0], nil
}
