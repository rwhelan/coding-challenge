package graph

import (
	"fmt"
)

type WalkerInstruction uint

const (
	PATH_CONTINUE WalkerInstruction = iota
	PATH_STOP
	PATH_DROP
	PATH_COPY
)

type contFunc func(p *Path, next *Node) WalkerInstruction

func Walk(g *Graph, start *Node, f contFunc) (*PathList, error) {
	pl := NewPathList()
	startNode := g.GetNode(start.Name)
	if startNode == nil {
		return nil, fmt.Errorf("start node %s not found in graph %s", start.Name, g.Name)
	}

	walk(pl, &Path{Nodes: []*Node{startNode}}, f)

	return pl, nil
}

func walkr(p *Path, f contFunc) *PathList {
	all := NewPathList()

	for _, e := range p.CurrentNode().Edges {
		switch f(p, e.Dst) {
		case PATH_DROP:
			continue

		case PATH_STOP:
			// Dedup
			if all.Len() == 0 ||
				all.Len() >= 1 &&
					!(all.Last() == p) {
				all.Append(p)
			}

		case PATH_COPY:
			all.Paths = append(all.Paths, p)
			fallthrough

		case PATH_CONTINUE:
			np := &Path{
				Nodes: append(p.Nodes, e.Dst),
				Edges: append(p.Edges, e),
				Cost:  p.Cost + e.Distance,
			}
			all.Add(walkr(np, f))
		}
	}

	return all
}

func walk(all *PathList, p *Path, f contFunc) {
	for _, e := range p.CurrentNode().Edges {
		switch f(p, e.Dst) {
		case PATH_DROP:
			continue

		case PATH_STOP:
			// Dedup
			if all.Len() == 0 ||
				all.Len() >= 1 &&
					!(all.Last() == p) {
				all.Append(p)
			}

		case PATH_COPY:
			all.Paths = append(all.Paths, p)
			fallthrough

		case PATH_CONTINUE:
			np := &Path{
				Nodes: append(p.Nodes, e.Dst),
				Edges: append(p.Edges, e),
				Cost:  p.Cost + e.Distance,
			}
			walk(all, np, f)

		}
	}
}
