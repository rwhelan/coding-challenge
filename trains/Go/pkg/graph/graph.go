package graph

import (
	"fmt"
	"strconv"
	"strings"
)

type Graph struct {
	Name  *string
	nodes map[string]*Node
}

func NewGraph(name string) *Graph {
	return &Graph{
		Name:  &name,
		nodes: make(map[string]*Node),
	}
}

func (g *Graph) AddNode(n *Node) {
	if _, exists := g.nodes[n.Name]; !exists {
		g.nodes[n.Name] = n
	}
}

func (g *Graph) GetNode(name string) *Node {
	if node, exists := g.nodes[name]; exists {
		return node
	}

	return nil
}

func (g *Graph) InitFromCommaString(s string) error {
	if len(g.nodes) != 0 {
		return fmt.Errorf("graph already initialized")
	}

	for _, t := range strings.Split(s, ",") {
		t = strings.Trim(t, " ")
		if len(t) != 3 {
			return fmt.Errorf("bad format; only single char node names supported")
		}

		src, dst, d := string(t[0]), string(t[1]), string(t[2])
		dis, err := strconv.Atoi(d)
		if err != nil {
			return fmt.Errorf("invalid distance %s for edge: %w", d, err)
		}

		if _, exists := g.nodes[src]; !exists {
			g.AddNode(NewNode(src))
		}

		if _, exists := g.nodes[dst]; !exists {
			g.AddNode(NewNode(dst))
		}

		g.GetNode(src).AddEdge(g.GetNode(dst), dis)
		//	fmt.Println(srcNode.Edges)
	}

	return nil
}

func (g *Graph) CalculatePath(route []string) (*Path, error) {
	p := &Path{}

	if len(route) == 0 {
		return p, nil
	}

	currentNode := g.GetNode(route[0])
	if currentNode == nil {
		return p, NewPathError(p, fmt.Sprintf("node %s not found in graph", route[0]))
	}

	p.Nodes = append(p.Nodes, currentNode)
	for _, nodeName := range route[1:] {
		edge, ok := currentNode.Edges[nodeName]
		if !ok {
			return p, NewPathError(
				p, fmt.Sprintf("node %s is not connected to node %s", p.CurrentNode().Name, nodeName),
			)
		}

		p.Nodes = append(p.Nodes, edge.Dst)
		p.Edges = append(p.Edges, edge)
		p.Cost += edge.Distance

		currentNode = edge.Dst
	}

	return p, nil

}

func (g *Graph) FindAllPaths(start, end *Node, min, max int) (*PathList, error) {
	walkFunction := func(p *Path, next *Node) WalkerInstruction {
		if len(p.Nodes) <= min+1 {
			return PATH_CONTINUE
		}

		if p.CurrentNode().Name == end.Name {
			return PATH_STOP
		}

		// The start node counts to the total - add 1
		if len(p.Nodes) >= max+1 {
			return PATH_DROP
		}

		return PATH_CONTINUE
	}

	if min > max {
		return nil, fmt.Errorf(fmt.Sprintf("min value of %d is larger than max value of %d", min, max))
	}

	startNode := g.GetNode(start.Name)
	endNode := g.GetNode(end.Name)

	if startNode == nil {
		return nil, fmt.Errorf("start node not found in graph")
	}

	if endNode == nil {
		return nil, fmt.Errorf("end node not found in graph")
	}

	pl := NewPathList()
	walk(pl, &Path{Nodes: []*Node{startNode}}, walkFunction)

	return pl, nil
}

func (g *Graph) FindShortestPath(start, end *Node, max int) (*Path, error) {
	walkFunction := func(p *Path, next *Node) WalkerInstruction {
		if len(p.Nodes) == 1 {
			return PATH_CONTINUE
		}

		if p.CurrentNode().Name == end.Name {
			return PATH_STOP
		}

		// The start node counts to the total - add 1
		if len(p.Nodes) >= max+1 {
			return PATH_DROP
		}

		return PATH_CONTINUE
	}

	startNode := g.GetNode(start.Name)
	endNode := g.GetNode(end.Name)

	if startNode == nil {
		return nil, fmt.Errorf("start node not found in graph")
	}

	if endNode == nil {
		return nil, fmt.Errorf("end node not found in graph")
	}

	pl := NewPathList()
	walk(pl, &Path{Nodes: []*Node{startNode}}, walkFunction)

	if len(pl.Paths) == 0 {
		return nil, fmt.Errorf("no path could be found between node %s and node %s", startNode.Name, endNode.Name)
	}

	path := &Path{}
	shortest := max

	for _, p := range pl.Paths {
		if p.Cost < shortest {
			shortest = p.Cost
			path = p
		}
	}

	return path, nil
}

func (g *Graph) FindAllPathsByDistance(start, end *Node, min, max int) (*PathList, error) {
	walkFunction := func(p *Path, next *Node) WalkerInstruction {
		if len(p.Nodes) == 1 || p.Cost < min {
			return PATH_CONTINUE
		}

		if p.Cost >= max {
			return PATH_DROP
		}

		if p.CurrentNode().Name == end.Name {
			return PATH_COPY
		}

		// if p.CurrentNode().Name == end.Name && p.Cost >= max {
		// 	return PATH_STOP

		// }
		// The start node counts to the total - add 1

		return PATH_CONTINUE
	}

	if min > max {
		return nil, fmt.Errorf(fmt.Sprintf("min value of %d is larger than max value of %d", min, max))
	}

	startNode := g.GetNode(start.Name)
	endNode := g.GetNode(end.Name)

	if startNode == nil {
		return nil, fmt.Errorf("start node not found in graph")
	}

	if endNode == nil {
		return nil, fmt.Errorf("end node not found in graph")
	}

	pl := NewPathList()
	walk(pl, &Path{Nodes: []*Node{startNode}}, walkFunction)

	pl.Dedup()

	return pl, nil
}
