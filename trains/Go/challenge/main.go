package main

import (
	"errors"
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

	// 1.) The distance of the route A-B-C.
	path, err := g.CalculatePath([]string{"A", "B", "C"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output #1: (9) %d\n", path.Cost)

	// 2.) The distance of the route A-D.
	path, err = g.CalculatePath([]string{"A", "D"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output #2: (5) %d\n", path.Cost)

	// 3.) The distance of the route A-D-C.
	path, err = g.CalculatePath([]string{"A", "D", "C"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output #3: (13) %d\n", path.Cost)

	// 4.) The distance of the route A-E-B-C-D.
	path, err = g.CalculatePath([]string{"A", "E", "B", "C", "D"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output #4: (22) %d\n", path.Cost)

	// 5.) The distance of the route A-E-D.
	if _, err = g.CalculatePath([]string{"A", "E", "C"}); err != nil {
		var pathError *graph.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Output #5: NO_SUCH_PATH")
		}
	} else {
		log.Fatal("#5 failed to return error")
	}

	// 6.) The number of trips starting at C and ending at C with a maximum
	//     of 3 stops. In the sample data below, there are two such trips:
	//     C-D-C (2 stops). and C-E-B-C (3 stops).
	startNode := g.GetNode("C")
	endNode := g.GetNode("C")
	pl, err := g.FindAllPaths(startNode, endNode, 0, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output #6: (2) %d\n", pl.Len())

	// 7.) The number of trips starting at A and ending at C with exactly 4
	//     stops. In the sample data below, there are three such trips:
	//     A to C (via B,C,D); A to C (via D,C,D); and A to C (via D,E,B).
	startNode = g.GetNode("A")
	endNode = g.GetNode("C")
	pl, err = g.FindAllPaths(startNode, endNode, 4, 4)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output #7: (3) %d\n", pl.Len())

	// 8.) The length of the shortest route (in terms of distance to travel)
	//     from A to C.
	startNode = g.GetNode("A")
	endNode = g.GetNode("C")
	path, err = g.FindShortestPath(startNode, endNode, 30)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output #8: (9) %d\n", path.Cost)

	// 9.) The length of the shortest route (in terms of distance to travel)
	//     from B to B.
	startNode = g.GetNode("B")
	endNode = g.GetNode("B")
	path, err = g.FindShortestPath(startNode, endNode, 30)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output #9: (9) %d\n", path.Cost)

	// 10.) The number of different routes from C to C with a distance of less
	//      than 30. In the sample data, the trips are:
	//      CDC, CEBC, CEBCDC, CDCEBC, CDEBC, CEBCEBC, CEBCEBCEBC.
	startNode = g.GetNode("C")
	endNode = g.GetNode("C")
	pl, err = g.FindAllPathsByDistance(startNode, endNode, 0, 30)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output #10: (7) %d\n", pl.Len())
}
