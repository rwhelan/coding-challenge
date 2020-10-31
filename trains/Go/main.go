package main

import (
	"fmt"
	"log"

	"github.com/rwhelan/coding-challenge/trains/Go/pkg/graph"
)

func main() {
	graphData := string("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7")
	g := graph.NewGraph("Trains")

	err := g.InitFromCommaString(graphData)
	if err != nil {
		log.Fatal(err)
	}

	// p, err := g.CalculatePath([]string{"A", "B", "C"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	startNode := g.GetNode("C")
	endNode := g.GetNode("C")
	pl, err := g.FindAllPathsByDistance(startNode, endNode, 0, 30)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range pl.Paths {
		fmt.Println(p)
	}
}

/////////////////////////////////////////////////////////////////
func SkipD(p *graph.Path, next *graph.Node) *bool {
	if len(p.Nodes) == 1 {
		return PathContinue()
	}
	if p.Cost >= 30 {
		return PathDrop()
	}

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
