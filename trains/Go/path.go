package main

import "fmt"

// type Path struct {
// 	Stops []*Town
// }

type OP uint

const (
	DROP OP = 1 << iota
	CONTINUE
	STOP
)

type contFunc func(p []*Town, current *Town) bool

var (
	allPaths [][]*Town
)

func walk(p []*Town, current *Town, f contFunc) {
	if f(p, current) {
		np := append(p, current)
		for _, r := range current.Routes {
			walk(np, r.Dst, f)
		}
	} else {
		allPaths = append(allPaths, p)
	}
}

func allps(p []*Town, current *Town) bool {
	if len(p) < 3 {
		return true
	}
	first := p[0]
	last := p[len(p)-1]

	fmt.Printf("FIRST: %s | LAST: %s | CURRNET: %s | PATH:",
		first.Name, last.Name, current.Name)
	printPath(p)

	fmt.Printf("FIRST: %p | LAST: %p | NEXT: %p\n",
		first, last, current)

	return first != last

}
