package graph

import (
	"testing"
)

func simpleGraph() *Graph {
	//
	//  A - 1 - B - 2 - C
	//           \
	//            3 - D - 4 - E - 5 - F
	//
	testGraph := NewGraph("TestGraph")

	nodeA := NewNode("A")
	nodeB := NewNode("B")
	nodeC := NewNode("C")
	nodeD := NewNode("D")
	nodeE := NewNode("E")
	nodeF := NewNode("F")

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

	return testGraph
}
func TestGraphInitFromCommaString(t *testing.T) {
	graphData := string("AB1, BC2, BD3")

	g := NewGraph("Test")
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
	testGraph := simpleGraph()

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

func TestGraphFindShortestPath(t *testing.T) {
	testGraph := simpleGraph()
	nodeA := testGraph.GetNode("A")
	nodeE := testGraph.GetNode("E")
	nodeF := testGraph.GetNode("F")

	nodeA.AddEdge(nodeE, 1)
	//    _________________
	//   /                 \
	//  A - 1 - B - 2 - C   1
	//           \           \
	//            3 - D - 4 - E - 5 - F
	//

	path, err := testGraph.FindShortestPath(nodeA, nodeF, 300)
	if err != nil {
		t.Fatal(err)
	}

	if path.Nodes[0].Name != "A" {
		t.Fatal("graph.FindShortestPath() invalid response: Node A")
	}

	if path.Nodes[1].Name != "E" {
		t.Fatal("graph.FindShortestPath() invalid response: Node E")
	}

	if path.Nodes[2].Name != "F" {
		t.Fatal("graph.FindShortestPath() invalid response: Node F")
	}

	if path.Nodes[0].Edges["E"].Dst != nodeE {
		t.Fatal("graph.FindShortestPath() invalid response: Edge A-E")
	}

	if path.Nodes[1].Edges["F"].Dst != nodeF {
		t.Fatal("graph.FindShortestPath() invalid response: Edge E-F")
	}

	if path.Nodes[0].Edges["E"].Distance != 1 {
		t.Fatal("graph.FindShortestPath() invalid response: Edge A-E Distance")
	}

	if path.Nodes[1].Edges["F"].Distance != 5 {
		t.Fatal("graph.FindShortestPath() invalid response: Edge E-F Distance")
	}
}

func TestGraphFindAllPaths(t *testing.T) {
	testGraph := simpleGraph()
	nodeA := testGraph.GetNode("A")
	nodeC := testGraph.GetNode("C")
	nodeF := testGraph.GetNode("F")

	nodeG := NewNode("G")
	testGraph.AddNode(nodeG)

	nodeC.AddEdge(nodeG, 6)
	nodeG.AddEdge(nodeF, 7)

	//
	//  A - 1 - B - 2 - C - 6 - G - 7
	//           \                   \
	//            3 - D - 4 - E - 5 - F
	//

	pl, err := testGraph.FindAllPaths(nodeA, nodeF, 1, 30)
	if err != nil {
		t.Fatal(err)
	}

	if pl.Len() != 2 {
		t.Fatal("graph.FindAllPaths() returns unexpected number of paths")
	}
}
