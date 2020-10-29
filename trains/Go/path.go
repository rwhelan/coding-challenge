package main

import "fmt"

type Path struct {
	Cost  int
	Stops []*Town
}

func (p *Path) CurrentTown() *Town {
	if len(p.Stops) == 0 {
		return nil
	}

	return p.Stops[len(p.Stops)-1]
}

// type OP uint

// const (
// 	DROP OP = 1 << iota
// 	CONTINUE
// 	STOP
// )

type contFunc func(p *Path, next *Town) bool

// var (
// 	allPaths [][]*Town
// )

func walk(allPaths *[]*Path, p *Path, f contFunc) {
	for _, r := range p.CurrentTown().Routes {
		if f(p, r.Dst) {
			np := &Path{
				Stops: append(p.Stops, r.Dst),
				Cost:  p.Cost + r.Distance,
			}
			walk(allPaths, np, f)
		} else {
			*allPaths = append(*allPaths, p)
		}
	}
}

func walkr(p *Path, f contFunc) []*Path {
	var allPaths []*Path

	for _, r := range p.CurrentTown().Routes {
		if f(p, r.Dst) {
			np := &Path{
				Stops: append(p.Stops, r.Dst),
				Cost:  p.Cost + r.Distance,
			}
			allPaths = append(allPaths, walkr(np, f)...)
		} else {
			allPaths = append(allPaths, p)
		}
	}

	return allPaths
}

func SkipD(p *Path, next *Town) bool {
	if pathContains(p, next) {
		fmt.Println("Currnet: ", p.CurrentTown().Name, "Next: ", next.Name)
		printPath(p)
		fmt.Println("FALSE - Contains\n")
		return false
	}
	if len(p.Stops) >= 10 {
		fmt.Println("Currnet: ", p.CurrentTown().Name, "Next: ", next.Name)
		printPath(p)
		fmt.Println("FALSE - Len\n")
		return false
	}

	// fmt.Println("TRUE\n")
	return true
}

// func DropLoopPaths(p *Path, next *Town) bool {
// 	if len(p.Stops) == 1 {
// 		return true
// 	}

// 	fmt.Println("CUrrnet: ", p.CurrentTown().Name, "Next: ", next.Name)
// 	printPath(p)
// 	return p.CurrentTown().Name != "C"
// }

func pathContains(p *Path, t *Town) bool {
	for _, i := range p.Stops {
		if i == t {
			return true
		}
	}

	return false
}

// func finalDst(t *Town) contFunc {
// 	return func(p *Path) []*Route {
// 		if len(p.Stops) > 15 {
// 			return nil
// 		}

// 		if p.CurrentNode() == t {
// 			return nil
// 		}

// 		return DropLoopPaths(p)
// 	}
// }

// //
// func DropLoopPaths(p *Path) []*Route {
// 	resp := p.CurrentNode().RouteList()

// 	for _, t := range p.Stops {
// 		for _, r := range resp {
// 			if r.Dst == t {
// 				resp = removeRoute(resp, r)
// 			}
// 		}
// 	}

// 	return resp
// }

// func removeRoute(allr []*Route, r *Route) []*Route {
// 	for i, cr := range allr {
// 		if cr == r {
// 			allr[i] = allr[len(allr)-1]
// 			return allr[:len(allr)-1]
// 		}
// 	}

// 	return allr
// }
