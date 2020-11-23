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

func (n *Node) AddEdge(on *Node, dis int) {
	e := &Edge{
		Src:      n,
		Dst:      on,
		Distance: dis,
	}

	n.Edges[on.Name] = e
}
