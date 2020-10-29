package main

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

func walk(p *Path, f contFunc) []*Path {
	var resp []*Path

	for _, r := range p.CurrentTown().Routes {
		if f(p, r.Dst) {
			np := &Path{
				Stops: append(p.Stops, r.Dst),
				Cost:  p.Cost + r.Distance,
			}
			resp = append(resp, walk(np, f)...)
		} else {
			resp = append(resp, p)
		}
	}

	return resp

	// if !f(p, next) {
	// 	return []*Path{
	// 		p,
	// 	}
	// } else {
	// 	if next != nil {
	// 		p.Stops = append(p.Stops, next)
	// 	}
	// 	resp := make([]*Path, 0)
	// 	fmt.Println(p.CurrentTown(), p.Stops)
	// 	for _, r := range p.CurrentTown().Routes {
	// 		if next != nil {
	// 			fmt.Printf("Currnet: %s Next: %s\n", p.CurrentTown().Name, r.Dst.Name)
	// 		}
	// 		resp = append(resp, walk(&Path{Stops: p.Stops, Cost: p.Cost + r.Distance}, r.Dst, f)...)
	// 	}
	// 	return resp
	// }
}

func DropLoopPaths(p *Path, next *Town) bool {
	if len(p.Stops) == 1 {
		return true
	}

	for _, t := range p.Stops[:len(p.Stops)-1] {
		if t == p.CurrentTown() {
			return false
		}
	}

	return true

}

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
