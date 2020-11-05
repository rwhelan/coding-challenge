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

func (g *Graph) Walk(start *Node, f contFunc) (*PathList, error) {
	startNode, ok := g.nodes[start.Name]
	if !ok {
		return nil, fmt.Errorf("start node %s not found in graph %s", (*start).Name, *g.Name)
	}

	pl := NewPathList()
	walk(pl, &Path{Nodes: []*Node{startNode}}, f)

	pl.Dedup()

	return pl, nil
}

func walkr(p *Path, f contFunc) *PathList {
	all := NewPathList()
	edges := p.CurrentNode().Edges
	if len(edges) == 0 {
		// leaf node
		inst := f(p, nil)
		if inst == PATH_COPY || inst == PATH_STOP {
			if all.Len() == 0 ||
				all.Len() >= 1 &&
					!(PathsEqual(all.Last(), p)) {
				all.Append(p.Duplicate())
			}
		}
	}
	for _, e := range edges {
		switch f(p, e.Dst) {
		case PATH_DROP:
			continue

		case PATH_STOP:
			// Dedup
			if all.Len() == 0 ||
				all.Len() >= 1 &&
					!(PathsEqual(all.Last(), p)) {
				all.Append(p.Duplicate())
			}

		case PATH_COPY:
			if all.Len() == 0 ||
				all.Len() >= 1 &&
					!(PathsEqual(all.Last(), p)) {

				all.Append(p.Duplicate())
			}
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
	edges := p.CurrentNode().Edges
	if len(edges) == 0 {
		// leaf node
		inst := f(p, nil)
		if inst == PATH_COPY || inst == PATH_STOP {
			if all.Len() == 0 ||
				all.Len() >= 1 &&
					!(PathsEqual(all.Last(), p)) {
				all.Append(p.Duplicate())
			}
		}
	}

	for _, e := range edges {
		switch f(p, e.Dst) {
		case PATH_DROP:
			continue

		case PATH_STOP:
			// Dedup
			if all.Len() == 0 ||
				all.Len() >= 1 &&
					!(PathsEqual(all.Last(), p)) {
				all.Append(p.Duplicate())
			}

		case PATH_COPY:
			if all.Len() == 0 ||
				all.Len() >= 1 &&
					!(PathsEqual(all.Last(), p)) {

				all.Append(p.Duplicate())
			}
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
