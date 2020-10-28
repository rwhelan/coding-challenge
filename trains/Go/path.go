package main

import "fmt"

type Path struct {
	Cost  int
	Stops []*Town
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

func allps(p *Path) []*Route {
	fmt.Println(p)
	if len(p.Stops) > 4 {
		return []*Route{}
	}

	return p.Stops[len(p.Stops)-1].RouteList()
}
