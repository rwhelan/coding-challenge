package main

import (
	"fmt"
	"log"

	"github.com/rwhelan/coding-challenge/trains/Go/pkg/graph"
)

// func graphLoader(input string, index map[string]*Node) {
// 	src, dst, dis := input[0], input[1], input[2]
// 	distance, _ := strconv.Atoi(string(dis))

// 	if _, ok := index[string(src)]; !ok {
// 		index[string(src)] = newNode(string(src))
// 	}

// 	if _, ok := index[string(dst)]; !ok {
// 		index[string(dst)] = newNode(string(dst))
// 	}

// 	newEdge(index[string(src)], index[string(dst)], distance)

// }

func main() {
	graphData := string("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7")
	g := graph.NewGraph("Trains")

	err := g.InitFromCommaString(graphData)
	if err != nil {
		log.Fatal(err)
	}

	// all := NewPathList()
	// walk(all, &Path{Nodes: []*Node{allNodes["A"]}}, SkipD)

	pathList, err := graph.Walk(g, g.GetNode("A"), SkipD)
	if err != nil {
		log.Fatal(err)
	}

	// for i, pth := range pathList.Paths {
	// 	fmt.Print(i, "   ")
	// 	fmt.Println(pth)
	// }

	fmt.Println(pathList)

}

func SkipD(p *graph.Path, next *graph.Node) *bool {
	if len(p.Nodes) == 1 {
		return PathContinue()
	}
	if p.Cost >= 30 {
		return PathDrop()
	}
	//fmt.Println(p.CurrentNode().Name, next.Name)
	//PrintPath(p)
	if p.CurrentNode().Name == "C" {
		return PathStop()
	}

	return PathContinue()
}

func PathContinue() *bool {
	t := true
	return &t
}

func PathStop() *bool {
	t := false
	return &t
}

func PathDrop() *bool {
	return nil
}
