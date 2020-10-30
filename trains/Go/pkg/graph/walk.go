package graph

import (
	"fmt"
)

type contFunc func(p *Path, next *Node) *bool

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
		cont := f(p, e.Dst)
		if cont == nil {
			continue
		}

		if *cont {
			np := &Path{
				Nodes: append(p.Nodes, e.Dst),
				Edges: append(p.Edges, e),
				Cost:  p.Cost + e.Distance,
			}
			all.Add(walkr(np, f))

		} else {
			// Dedup
			if all.Len() == 0 ||
				all.Len() >= 1 &&
					!(all.Last() == p) {
				all.Append(p)
			}
		}
	}

	return all
}

func walk(all *PathList, p *Path, f contFunc) {
	for _, e := range p.CurrentNode().Edges {
		cont := f(p, e.Dst)
		if cont == nil {
			continue
		}

		if *cont {
			np := &Path{
				Nodes: append(p.Nodes, e.Dst),
				Edges: append(p.Edges, e),
				Cost:  p.Cost + e.Distance,
			}
			walk(all, np, f)

		} else {
			// Dedup
			if all.Len() == 0 ||
				all.Len() >= 1 &&
					!(all.Last() == p) {
				all.Append(p)
			}
		}
	}
}
