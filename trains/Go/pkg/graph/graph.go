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
