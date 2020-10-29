package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Town struct {
	Name   string
	Routes map[string]*Route
}

func (t *Town) RouteList() []*Route {
	routes := make([]*Route, 0, len(t.Routes))
	for _, r := range t.Routes {
		routes = append(routes, r)
	}

	return routes
}

type Route struct {
	Src      *Town
	Dst      *Town
	Distance int
}

func newTown(name string) *Town {
	return &Town{
		Name:   name,
		Routes: make(map[string]*Route),
	}
}

func newRoute(src, dst *Town, dist int) {
	src.Routes[dst.Name] = &Route{
		Src:      src,
		Dst:      dst,
		Distance: dist,
	}
}

func graphLoader(input string, index map[string]*Town) {
	src, dst, dis := input[0], input[1], input[2]
	distance, _ := strconv.Atoi(string(dis))

	if _, ok := index[string(src)]; !ok {
		index[string(src)] = newTown(string(src))
	}

	if _, ok := index[string(dst)]; !ok {
		index[string(dst)] = newTown(string(dst))
	}

	newRoute(index[string(src)], index[string(dst)], distance)

}

func main() {
	allTowns := make(map[string]*Town)

	graphData := string("AB5, BC4, CD8, DC8, DE6, AD5, CE2, EB3, AE7")

	for _, edge := range strings.Split(graphData, ",") {
		edge := strings.Trim(edge, " ")
		graphLoader(edge, allTowns)
	}

	allPaths := make([]*Path, 0)
	walk(
		&allPaths, &Path{Stops: []*Town{allTowns["A"]}}, SkipD,
	)

	allp := walkr(&Path{Stops: []*Town{allTowns["A"]}}, SkipD)

	// allPaths := walk(
	// 	&Path{Stops: []*Town{allTowns["A"]}}, SkipD,
	// )

	fmt.Println("ALL: ", allPaths)
	for i, pth := range allPaths {
		fmt.Print(i)
		printPath(pth)
	}

	fmt.Println("AAL P: ", allp)
	for i, pth := range allp {
		fmt.Print(i)
		printPath(pth)
	}
}

func printPath(path *Path) {
	for _, t := range path.Stops {
		fmt.Printf(" => %s", (*t).Name)
	}

	fmt.Printf("   Cost: %d\n", path.Cost)
}
