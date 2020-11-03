package tests

import (
	"testing"

	"github.com/rwhelan/coding-challenge/trains/Go/pkg/graph"
)

func TestGraphInitFromCommaString(t *testing.T) {
	graphData := string("AB1, BC2, BD3")

	g := graph.NewGraph("Test")
	err := g.InitFromCommaString(graphData)
	if err != nil {
		t.Fatal(err)
	}

	err = g.InitFromCommaString(graphData)
	if err == nil {
		t.Fatal("graph.InitFromCommaString() did not error on re-init")
	}

	nodeA := g.GetNode("A")
	if nodeA == nil {
		t.Fatal("graph.GetNode(\"A\") failed to return expected node")
	}

	nodeB := g.GetNode("B")
	if nodeB == nil {
		t.Fatal("graph.GetNode(\"B\") failed to return expected node")
	}

	nodeC := g.GetNode("C")
	if nodeC == nil {
		t.Fatal("graph.GetNode(\"C\") failed to return expected node")
	}

	nodeD := g.GetNode("D")
	if nodeD == nil {
		t.Fatal("graph.GetNode(\"D\") failed to return expected node")
	}

	if nodeA.Edges["B"].Dst != nodeB {
		t.Fatal("graph.InitFromCommaString() failed to construct expected graph: Edge A to B")
	}

	if nodeA.Edges["B"].Distance != 1 {
		t.Fatal("graph.InitFromCommaString() failed to construct expected graph: Edge A to B distance")
	}

	if nodeB.Edges["C"].Dst != nodeC {
		t.Fatal("graph.InitFromCommaString() failed to construct expected graph: Edge B to C")
	}

	if nodeB.Edges["C"].Distance != 2 {
		t.Fatal("graph.InitFromCommaString() failed to construct expected graph: Edge B to C distance")
	}

	if nodeB.Edges["D"].Dst != nodeD {
		t.Fatal("graph.InitFromCommaString() failed to construct expected graph: Edge B to D")
	}

	if nodeB.Edges["D"].Distance != 3 {
		t.Fatal("graph.InitFromCommaString() failed to construct expected graph: Edge B to D distance")
	}
}

func TestGraphCalculatePath(t *testing.T) {
	testGraph := graph.NewGraph("TestGraph")

	nodeA := graph.NewNode("A")
	nodeB := graph.NewNode("B")
	nodeC := graph.NewNode("C")
	nodeD := graph.NewNode("D")
	nodeE := graph.NewNode("E")
	nodeF := graph.NewNode("F")

	testGraph.AddNode(nodeA)
	testGraph.AddNode(nodeB)
	testGraph.AddNode(nodeC)
	testGraph.AddNode(nodeD)
	testGraph.AddNode(nodeE)
	testGraph.AddNode(nodeF)

	nodeA.AddEdge(nodeB, 1)
	nodeB.AddEdge(nodeC, 2)
	nodeB.AddEdge(nodeD, 3)
	nodeD.AddEdge(nodeE, 4)
	nodeE.AddEdge(nodeF, 5)

	path, err := testGraph.CalculatePath([]string{"A", "B", "D", "E", "F"})
	if err != nil {
		t.Fatal(err)
	}

	if path.Cost != 13 {
		t.Fatal("graph.CalculatePath() bad cost value")
	}

	_, err = testGraph.CalculatePath([]string{"A", "B", "D", "E", "F", "A"})
	if err == nil {
		t.Fatal("graph.CalculatePath() expected err on bad path")
	}
}
