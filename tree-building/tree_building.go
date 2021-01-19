package tree

import (
	"errors"
	"sort"
)

// Record represent a flat post element
type Record struct {
	ID     int
	Parent int
}

// Node represent an element of a tree structure
type Node struct {
	ID       int
	Children []*Node
}

// Build construct a Tree for a given array of records
func Build(records []Record) (*Node, error) {
	nodePosition := make(map[int]*Node, len(records))
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	for i, record := range records {
		if record.ID != i || record.Parent > record.ID || record.ID > 0 && record.Parent == record.ID {
			return nil, errors.New("not in sequence or has bad parent")
		}

		nodePosition[record.ID] = &Node{ID: record.ID}
		parent := nodePosition[record.Parent]
		n := &Node{ID: record.ID}
		parent.Children = append(parent.Children, n)
		nodePosition[n.ID] = n
	}

	return nodePosition[0], nil
}
