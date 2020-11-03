package tests

import "github.com/rwhelan/coding-challenge/trains/Go/pkg/graph"

func generateSimplePath() *graph.Path {
	// One - 2 - Two - 4 - Three
	nOne := graph.NewNode("One")
	nTwo := graph.NewNode("Two")
	nThree := graph.NewNode("Three")

	eOne := &graph.Edge{
		Src:      nOne,
		Dst:      nTwo,
		Distance: 2,
	}

	eTwo := &graph.Edge{
		Src:      nTwo,
		Dst:      nThree,
		Distance: 4,
	}

	nOne.Edges["Two"] = eOne
	nTwo.Edges["Three"] = eTwo

	testPath := &graph.Path{
		Nodes: []*graph.Node{
			nOne, nTwo, nThree,
		},
		Edges: []*graph.Edge{
			eOne, eTwo,
		},
	}

	for _, e := range testPath.Edges {
		testPath.Cost += e.Distance
	}

	return testPath
}
