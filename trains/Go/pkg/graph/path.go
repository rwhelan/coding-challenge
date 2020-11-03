package graph

import (
	"fmt"
	"strconv"
	"strings"
)

type Path struct {
	Cost  int
	Nodes []*Node
	Edges []*Edge
}

func (p *Path) String() string {
	s := strings.Builder{}

	for i := 0; i < len(p.Nodes); i++ {
		s.WriteString(p.Nodes[i].Name)
		if len(p.Nodes)-1 != i {
			s.WriteString("==(")
		}
		if (len(p.Edges) - 1) >= i {
			s.WriteString(strconv.Itoa(p.Edges[i].Distance))
			s.WriteString(")==>")
		}
	}

	s.WriteString(":")
	s.WriteString(strconv.Itoa(p.Cost))
	return s.String()
}

func (p *Path) CurrentNode() *Node {
	if len(p.Nodes) == 0 {
		return nil
	}

	return p.Nodes[len(p.Nodes)-1]
}

type PathList struct {
	Paths []*Path
}

func NewPathList() *PathList {
	return &PathList{
		Paths: []*Path{},
	}
}
func (p *PathList) Len() int {
	return len(p.Paths)
}

func (p *PathList) Last() *Path {
	if len(p.Paths) == 0 {
		return nil
	}

	return p.Paths[len(p.Paths)-1]
}

func (p *PathList) Add(lp *PathList) {
	for _, path := range lp.Paths {
		p.Append(path)
	}
}

func (p *PathList) Append(path ...*Path) {
	p.Paths = append(p.Paths, path...)
}

func (pl *PathList) Dedup() {
	for i, p := range pl.Paths {
		if p == nil {
			continue
		}
		for x, y := range pl.Paths[i+1:] {
			if y == nil {
				continue
			}
			n := x + i + 1
			if PathsEqual(p, y) {
				pl.Paths[n] = nil
			}
		}
	}

	n := make([]*Path, 0)
	for _, p := range pl.Paths {
		if p != nil {
			n = append(n, p)
		}
	}

	pl.Paths = n
}

func PathsEqual(x, y *Path) bool {
	if len(x.Nodes) != len(y.Nodes) {
		return false
	}

	for i := 0; i < len(x.Nodes)-1; i++ {
		if x.Nodes[i] != y.Nodes[i] {
			return false
		}
	}

	// Assert Path Costs
	if x.Cost != y.Cost {
		fmt.Println(x)
		fmt.Println(y)
		panic("COSTS SHOULD MATCH")
	}

	return true
}

func PathContains(p *Path, t *Node) bool {
	for _, i := range p.Nodes {
		if i == t {
			return true
		}
	}

	return false
}

// func PathDedup(allPaths []*Path) []*Path {
// 	for i, p := range allPaths {
// 		if p == nil {
// 			continue
// 		}
// 		for x, y := range allPaths[i+1:] {
// 			if y == nil {
// 				continue
// 			}
// 			n := x + i + 1
// 			if PathsEqual(p, y) {
// 				allPaths[n] = nil
// 			}
// 		}
// 	}

// 	n := make([]*Path, 0)
// 	for _, p := range allPaths {
// 		if p != nil {
// 			n = append(n, p)
// 		}
// 	}
// 	return n
// }

// func PrintPath(path *Path) {
// 	s := strings.Builder{}

// 	for i := 0; i < len(path.Nodes); i++ {
// 		s.WriteString(path.Nodes[i].Name)
// 		if len(path.Nodes)-1 != i {
// 			s.WriteString(" ==(")
// 		}
// 		if (len(path.Edges) - 1) >= i {
// 			s.WriteString(strconv.Itoa(path.Edges[i].Distance))
// 			s.WriteString(")==> ")
// 		}
// 	}

// 	s.WriteString(" : ")
// 	s.WriteString(strconv.Itoa(path.Cost))
// 	fmt.Println(s.String())
// }
