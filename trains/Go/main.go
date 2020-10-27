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

func parser(input string, index map[string]*Town) {
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
		parser(edge, allTowns)
	}

	for k, v := range allTowns {
		fmt.Printf("%s\n", k)
		for _, t := range v.Routes {
			fmt.Println("    ", t.Dst.Name)
		}
	}
}
