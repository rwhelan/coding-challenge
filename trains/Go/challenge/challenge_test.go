package challenge

import (
	"errors"
	"fmt"
	"testing"

	"github.com/rwhelan/coding-challenge/trains/Go/pkg/graph"
)

const (
	graphData = "AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7"
)

// 1.) The distance of the route A-B-C.
func TestNumberOne(t *testing.T) {
	g := graph.NewGraph("One")
	err := g.InitFromCommaString(graphData)
	if err != nil {
		t.Fatal(err)
	}

	path, err := g.CalculatePath([]string{"A", "B", "C"})
	if err != nil {
		t.Fatal(err)
	}

	if path.Cost != 9 {
		t.Fatalf("Test One: expected 9, got %d", path.Cost)
	}

	fmt.Printf("Output #1: %d\n", path.Cost)
}

// 2.) The distance of the route A-D.
func TestNumberTwo(t *testing.T) {
	g := graph.NewGraph("Two")
	err := g.InitFromCommaString(graphData)
	if err != nil {
		t.Fatal(err)
	}

	path, err := g.CalculatePath([]string{"A", "D"})
	if err != nil {
		t.Fatal(err)
	}

	if path.Cost != 5 {
		t.Fatalf("Test Two: expected 5, got %d", path.Cost)
	}

	fmt.Printf("Output #2: %d\n", path.Cost)
}

// 3.) The distance of the route A-D-C.
func TestNumberThree(t *testing.T) {
	g := graph.NewGraph("Three")
	err := g.InitFromCommaString(graphData)
	if err != nil {
		t.Fatal(err)
	}

	path, err := g.CalculatePath([]string{"A", "D", "C"})
	if err != nil {
		t.Fatal(err)
	}

	if path.Cost != 13 {
		t.Fatalf("Test Three: expected 13, got %d", path.Cost)
	}

	fmt.Printf("Output #3: %d\n", path.Cost)
}

// 4.) The distance of the route A-E-B-C-D.
func TestNumberFour(t *testing.T) {
	g := graph.NewGraph("Four")
	err := g.InitFromCommaString(graphData)
	if err != nil {
		t.Fatal(err)
	}

	path, err := g.CalculatePath([]string{"A", "E", "B", "C", "D"})
	if err != nil {
		t.Fatal(err)
	}

	if path.Cost != 22 {
		t.Fatalf("Test Four: expected 22, got %d", path.Cost)
	}

	fmt.Printf("Output #4: %d\n", path.Cost)
}

// 5.) The distance of the route A-E-D. (No Such Path)
func TestNumberFive(t *testing.T) {
	g := graph.NewGraph("Five")
	err := g.InitFromCommaString(graphData)
	if err != nil {
		t.Fatal(err)
	}

	_, err = g.CalculatePath([]string{"A", "E", "D"})
	if err == nil {
		t.Fatal("Test Five: expected Error")
	}

	var pathError *graph.PathError
	if errors.As(err, &pathError) {
		fmt.Println("Output #5: NO_SUCH_PATH")
	} else {
		t.Fatal(err)
	}
}

// 6.) The number of trips starting at C and ending at C with a maximum of 3 stops.
//     In the sample data, there are two such trips: C-D-C (2 stops). and C-E-B-C (3 stops).
func TestNumberSix(t *testing.T) {
	g := graph.NewGraph("Six")
	err := g.InitFromCommaString(graphData)
	if err != nil {
		t.Fatal(err)
	}

	f := func(p *graph.Path, next *graph.Node) graph.WalkerInstruction {
		if len(p.Nodes) == 1 {
			return graph.PATH_CONTINUE
		}

		if len(p.Nodes) >= 5 {
			return graph.PATH_DROP
		}

		if p.CurrentNode().Name == "C" {
			return graph.PATH_STOP
		}

		return graph.PATH_CONTINUE
	}

	pathList, err := g.Walk(g.GetNode("C"), f)
	if err != nil {
		t.Fatal(err)
	}

	if pathList.Len() != 2 {
		t.Fatalf("Test Six: expected 2 paths, got %d\n", pathList.Len())
	}

	fmt.Printf("Output #6: %d\n", pathList.Len())
}

// 7.) The number of trips starting at A and ending at C with exactly 4 stops.
//     In the sample data, there are three such trips:
//     A to C (via B,C,D); A to C (via D,C,D); and A to C (via D,E,B).
func TestNumberSeven(t *testing.T) {
	g := graph.NewGraph("Seven")
	err := g.InitFromCommaString(graphData)
	if err != nil {
		t.Fatal(err)
	}

	f := func(p *graph.Path, next *graph.Node) graph.WalkerInstruction {
		if len(p.Nodes) <= 4 {
			return graph.PATH_CONTINUE
		}

		if len(p.Nodes) > 5 {
			return graph.PATH_DROP
		}

		if p.CurrentNode().Name == "C" {
			return graph.PATH_STOP
		}

		return graph.PATH_CONTINUE
	}

	pathList, err := g.Walk(g.GetNode("A"), f)
	if err != nil {
		t.Fatal(err)
	}

	if pathList.Len() != 3 {
		t.Fatalf("Test Seven: expected 3 paths, got %d\n", pathList.Len())
	}

	fmt.Printf("Output #7: %d\n", pathList.Len())
}

// 8.) The length of the shortest route (in terms of distance to travel) from A to C.
func TestNumberEight(t *testing.T) {
	g := graph.NewGraph("Eight")
	err := g.InitFromCommaString(graphData)
	if err != nil {
		t.Fatal(err)
	}

	path, err := g.FindShortestPath(g.GetNode("A"), g.GetNode("C"), 30)
	if err != nil {
		t.Fatal(err)
	}

	if path.Cost != 9 {
		t.Fatalf("Test Eight: expected path distance of 9, got %d\n", path.Cost)
	}

	fmt.Printf("Output #8: %d\n", path.Cost)
}

// 9.) The length of the shortest route (in terms of distance to travel) from B to B.
func TestNumberNine(t *testing.T) {
	g := graph.NewGraph("Nine")
	err := g.InitFromCommaString(graphData)
	if err != nil {
		t.Fatal(err)
	}

	path, err := g.FindShortestPath(g.GetNode("B"), g.GetNode("B"), 30)
	if err != nil {
		t.Fatal(err)
	}

	if path.Cost != 9 {
		t.Fatalf("Test Nine: expected path distance of 9, got %d\n", path.Cost)
	}

	fmt.Printf("Output #9: %d\n", path.Cost)
}

// 10.) The number of different routes from C to C with a distance of less than 30.
//      In the sample, the trips are:
//      CDC, CEBC, CEBCDC, CDCEBC, CDEBC, CEBCEBC, CEBCEBCEBC.
func TestNumberTen(t *testing.T) {
	g := graph.NewGraph("Ten")
	err := g.InitFromCommaString(graphData)
	if err != nil {
		t.Fatal(err)
	}

	f := func(p *graph.Path, next *graph.Node) graph.WalkerInstruction {
		if len(p.Nodes) == 1 {
			return graph.PATH_CONTINUE
		}

		if p.Cost >= 30 {
			return graph.PATH_DROP
		}

		if p.CurrentNode().Name == "C" {
			return graph.PATH_COPY
		}

		return graph.PATH_CONTINUE
	}

	pathList, err := g.Walk(g.GetNode("C"), f)
	if err != nil {
		t.Fatal(err)
	}

	if pathList.Len() != 7 {
		t.Fatalf("Test Ten: expected 7 paths, got %d\n", pathList.Len())
	}

	fmt.Printf("Output #10: %d\n", pathList.Len())
}
