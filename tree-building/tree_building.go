package tree

import (
	"errors"
	"sort"
)

type Record struct{
	ID int
	Parent int
}
type Node struct {
	ID int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	treeMap := make(map[int]*Node, len(records))
	for i, v := range records {
		if v.ID != i || v.Parent > v.ID || v.ID > 0 && v.Parent == v.ID {
			return nil, errors.New("error")
		}
		newNode := &Node{}
		newNode.ID = v.ID
		treeMap[v.ID] = newNode
		if v.ID != 0 {
			parentNode := treeMap[v.Parent]
			parentNode.Children = append(parentNode.Children, newNode)
		}
	}
	return treeMap[0], nil
}