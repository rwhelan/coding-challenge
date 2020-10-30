package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Name  string
	Edges map[string]*Edge
}

func (t *Node) EdgeList() []*Edge {
	Edges := make([]*Edge, 0, len(t.Edges))
	for _, r := range t.Edges {
		Edges = append(Edges, r)
	}

	return Edges
}

type Edge struct {
	Src      *Node
	Dst      *Node
	Distance int
}

func newNode(name string) *Node {
	return &Node{
		Name:  name,
		Edges: make(map[string]*Edge),
	}
}

func newEdge(src, dst *Node, dist int) {
	src.Edges[dst.Name] = &Edge{
		Src:      src,
		Dst:      dst,
		Distance: dist,
	}
}

func graphLoader(input string, index map[string]*Node) {
	src, dst, dis := input[0], input[1], input[2]
	distance, _ := strconv.Atoi(string(dis))

	if _, ok := index[string(src)]; !ok {
		index[string(src)] = newNode(string(src))
	}

	if _, ok := index[string(dst)]; !ok {
		index[string(dst)] = newNode(string(dst))
	}

	newEdge(index[string(src)], index[string(dst)], distance)

}

func main() {
	allNodes := make(map[string]*Node)

	graphData := string("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7")

	for _, edge := range strings.Split(graphData, ",") {
		edge := strings.Trim(edge, " ")
		graphLoader(edge, allNodes)
	}

	// all := NewPathList()
	// walk(all, &Path{Nodes: []*Node{allNodes["A"]}}, SkipD)

	all := walkr(&Path{Nodes: []*Node{allNodes["A"]}}, SkipD)

	for i, pth := range all.Paths {
		fmt.Print(i, "   ")
		fmt.Println(pth)
	}

	// np := PathDedup(allp)
	fmt.Println(all)
	//fmt.Println(np)

	// fmt.Println("AAL ddup: ", np)
	// for i, pth := range np {
	// 	fmt.Print(i)
	// 	printPath(pth)
	// }
}
