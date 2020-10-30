package main

import "fmt"

func PathContains(p *Path, t *Town) bool {
	for _, i := range p.Stops {
		if i == t {
			return true
		}
	}

	return false
}

func PathDedup(allPaths []*Path) []*Path {
	for i, p := range allPaths {
		if p == nil {
			continue
		}
		for x, y := range allPaths[i+1:] {
			if y == nil {
				continue
			}
			n := x + i + 1
			if PathsEqual(p, y) {
				allPaths[n] = nil
			}
		}
	}

	n := make([]*Path, 0)
	for _, p := range allPaths {
		if p != nil {
			n = append(n, p)
		}
	}
	return n
}

func PathsEqual(x, y *Path) bool {
	if len(x.Stops) != len(y.Stops) {
		return false
	}

	for i := 0; i < len(x.Stops)-1; i++ {
		if x.Stops[i] != y.Stops[i] {
			return false
		}
	}

	// Assert Path Costs
	if x.Cost != y.Cost {
		PrintPath(x)
		PrintPath(y)
		panic("COSTS SHOULD MATCH")
	}

	return true
}

func PrintPath(path *Path) {
	for _, t := range path.Stops {
		fmt.Printf(" => %s", (*t).Name)
	}

	fmt.Printf("   Cost: %d\n", path.Cost)
}
