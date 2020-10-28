package main

import "fmt"

// type Path struct {
// 	Stops []*Town
// }

// type OP uint

// const (
// 	DROP OP = 1 << iota
// 	CONTINUE
// 	STOP
// )

type contFunc func(p []*Town) []*Route

// var (
// 	allPaths [][]*Town
// )

func walk(p []*Town, f contFunc) [][]*Town {
	cont := f(p)

	if len(cont) == 0 {
		return [][]*Town{
			p,
		}
	}

	resp := make([][]*Town, 0)
	for _, r := range cont {
		np := append(p, r.Dst)
		resp = append(resp, walk(np, f)...)
	}

	return resp
}

func allps(p []*Town) []*Route {
	fmt.Println(p)
	if len(p) > 4 {
		return []*Route{}
	}

	return p[len(p)-1].RouteList()
}
