package main

type Path struct {
	Cost  int
	Stops []*Town
}

func (p *Path) CurrentNode() *Town {
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

type contFunc func(p *Path) []*Route

// var (
// 	allPaths [][]*Town
// )

func walk(p *Path, f contFunc) []*Path {
	cont := f(p)

	if len(cont) == 0 {
		return []*Path{
			p,
		}
	}

	resp := make([]*Path, 0)
	for _, r := range cont {
		np := append(p.Stops, r.Dst)
		resp = append(resp, walk(&Path{Stops: np, Cost: p.Cost + r.Distance}, f)...)
	}

	return resp
}

func finalDst(t *Town) contFunc {
	return func(p *Path) []*Route {
		if len(p.Stops) > 15 {
			return nil
		}

		if p.CurrentNode() == t {
			return nil
		}

		return DropLoopPaths(p)
	}
}

//
func DropLoopPaths(p *Path) []*Route {
	resp := p.CurrentNode().RouteList()

	for _, t := range p.Stops {
		for _, r := range resp {
			if r.Dst == t {
				resp = removeRoute(resp, r)
			}
		}
	}

	return resp
}

func removeRoute(allr []*Route, r *Route) []*Route {
	for i, cr := range allr {
		if cr == r {
			allr[i] = allr[len(allr)-1]
			return allr[:len(allr)-1]
		}
	}

	return allr
}
