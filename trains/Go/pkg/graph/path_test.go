package graph

import (
	"testing"
)

func generateSimplePath() *Path {
	// One - 2 - Two - 4 - Three
	nOne := NewNode("One")
	nTwo := NewNode("Two")
	nThree := NewNode("Three")

	eOne := &Edge{
		Src:      nOne,
		Dst:      nTwo,
		Distance: 2,
	}

	eTwo := &Edge{
		Src:      nTwo,
		Dst:      nThree,
		Distance: 4,
	}

	nOne.Edges["Two"] = eOne
	nTwo.Edges["Three"] = eTwo

	testPath := &Path{
		Nodes: []*Node{
			nOne, nTwo, nThree,
		},
		Edges: []*Edge{
			eOne, eTwo,
		},
	}

	for _, e := range testPath.Edges {
		testPath.Cost += e.Distance
	}

	return testPath
}

func TestPathCurrentNode(t *testing.T) {
	nOne := NewNode("One")
	nTwo := NewNode("Two")
	nThree := NewNode("Three")

	testPath := &Path{
		Nodes: []*Node{
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
	test2Path := &Path{
		Cost:  testPath.Cost,
		Nodes: make([]*Node, len(testPath.Nodes)),
		Edges: make([]*Edge, len(testPath.Edges)),
	}

	copy(test2Path.Nodes, testPath.Nodes)
	copy(test2Path.Edges, testPath.Edges)

	if !PathsEqual(testPath, test2Path) {
		t.Fatal("PathEqual() failed equal")
	}

	if PathsEqual(testPath, generateSimplePath()) {
		t.Fatal("PathEqual() non-equal equal")
	}
}

func TestPathContains(t *testing.T) {
	testPath := generateSimplePath()
	testNode := testPath.Nodes[1]

	if !PathContains(testPath, testNode) {
		t.Fatal("PathContains() failed to match")
	}

	if PathContains(testPath, NewNode("Three")) {
		t.Fatal("PathContains() found non-match")
	}
}
