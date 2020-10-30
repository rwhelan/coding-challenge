package graph

type Node struct {
	Name  string
	Edges map[string]*Edge
}

func NewNode(name string) *Node {
	return &Node{
		Name:  name,
		Edges: make(map[string]*Edge),
	}
}

func (n *Node) EdgeList() []*Edge {
	Edges := make([]*Edge, 0, len(n.Edges))
	for _, r := range n.Edges {
		Edges = append(Edges, r)
	}

	return Edges
}

func (n *Node) AddEdge(on *Node, dis int) {
	e := &Edge{
		Src:      n,
		Dst:      on,
		Distance: dis,
	}

	n.Edges[on.Name] = e
}
