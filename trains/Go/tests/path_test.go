package tests

import (
	"testing"

	"github.com/rwhelan/coding-challenge/trains/Go/pkg/graph"
)

func TestPathCurrentNode(t *testing.T) {
	nOne := graph.NewNode("One")
	nTwo := graph.NewNode("Two")
	nThree := graph.NewNode("Three")

	testPath := &graph.Path{
		Nodes: []*graph.Node{
			nOne, nTwo, nThree,
		},
	}

	if testPath.CurrentNode() != nThree {
		t.Fatal("CurrentNode() unexpected returned ")
	}

	if testPath.CurrentNode().Name != "Three" {
		t.Fatal("CurrentNode() unexpected .Name on returned Node")
	}
}

func TestPathDuplicate(t *testing.T) {
	testPath := generateSimplePath()

	duplicatedPath := testPath.Duplicate()

	for i := 0; i < len(testPath.Nodes); i++ {
		if testPath.Nodes[i] != duplicatedPath.Nodes[i] {
			t.Fatalf("path.Duplicate() node %s != %s",
				testPath.Nodes[i].Name, duplicatedPath.Nodes[i].Name,
			)
		}
	}

	for i := 0; i < len(testPath.Edges); i++ {
		if testPath.Edges[i] != duplicatedPath.Edges[i] {
			t.Fatalf("path.Duplicate() edge %p != %p",
				testPath.Edges[i], duplicatedPath.Edges[i],
			)
		}
	}

	if testPath.Cost != duplicatedPath.Cost {
		t.Fatal("path.Duplicate() .Cost not equal")
	}
}

func TestPathEqual(t *testing.T) {
	testPath := generateSimplePath()
	test2Path := &graph.Path{
		Cost:  testPath.Cost,
		Nodes: make([]*graph.Node, len(testPath.Nodes)),
		Edges: make([]*graph.Edge, len(testPath.Edges)),
	}

	copy(test2Path.Nodes, testPath.Nodes)
	copy(test2Path.Edges, testPath.Edges)

	if !graph.PathsEqual(testPath, test2Path) {
		t.Fatal("PathEqual() failed equal")
	}

	if graph.PathsEqual(testPath, generateSimplePath()) {
		t.Fatal("PathEqual() non-equal equal")
	}
}

func TestPathContains(t *testing.T) {
	testPath := generateSimplePath()
	testNode := testPath.Nodes[1]

	if !graph.PathContains(testPath, testNode) {
		t.Fatal("PathContains() failed to match")
	}

	if graph.PathContains(testPath, graph.NewNode("Three")) {
		t.Fatal("PathContains() found non-match")
	}
}
