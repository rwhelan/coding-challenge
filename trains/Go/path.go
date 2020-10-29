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

// type contFunc func(p *Path, next *Town) bool

// func walk(allPaths *[]*Path, p *Path, f contFunc) {
// 	for _, r := range p.CurrentTown().Routes {
// 		if f(p, r.Dst) {
// 			np := &Path{
// 				Stops: append(p.Stops, r.Dst),
// 				Cost:  p.Cost + r.Distance,
// 			}
// 			walk(allPaths, np, f)
// 		} else {
// 			*allPaths = append(*allPaths, p)
// 		}
// 	}
// }

// func walkr(p *Path, f contFunc) []*Path {
// 	var allPaths []*Path

// 	for _, r := range p.CurrentTown().Routes {
// 		if f(p, r.Dst) {
// 			np := &Path{
// 				Stops: append(p.Stops, r.Dst),
// 				Cost:  p.Cost + r.Distance,
// 			}
// 			allPaths = append(allPaths, walkr(np, f)...)
// 		} else {
// 			allPaths = append(allPaths, p)
// 		}
// 	}

// 	return allPaths
// }

// func SkipD(p *Path, next *Town) bool {
// 	if pathContains(p, next) {
// 		fmt.Println("Currnet: ", p.CurrentTown().Name, "Next: ", next.Name)
// 		printPath(p)
// 		fmt.Println("FALSE - Contains\n")
// 		return false
// 	}
// 	if len(p.Stops) >= 10 {
// 		fmt.Println("Currnet: ", p.CurrentTown().Name, "Next: ", next.Name)
// 		printPath(p)
// 		fmt.Println("FALSE - Len\n")
// 		return false
// 	}

// 	return true
// }

func pathContains(p *Path, t *Town) bool {
	for _, i := range p.Stops {
		if i == t {
			return true
		}
	}

	return false
}

type contFunc func(p *Path, next *Town) *bool

// func walk(allPaths *[]*Path, p *Path, f contFunc) {
// 	for _, r := range p.CurrentTown().Routes {
// 		if f(p, r.Dst) {
// 			np := &Path{
// 				Stops: append(p.Stops, r.Dst),
// 				Cost:  p.Cost + r.Distance,
// 			}
// 			walk(allPaths, np, f)
// 		} else {
// 			*allPaths = append(*allPaths, p)
// 		}
// 	}
// }

func walkr(p *Path, f contFunc) []*Path {
	allPaths := []*Path{}

	for _, r := range p.CurrentTown().Routes {
		cont := f(p, r.Dst)
		if cont == nil {
			continue
		}

		if *cont {
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

func walk(allPaths *[]*Path, p *Path, f contFunc) {
	for _, r := range p.CurrentTown().Routes {
		cont := f(p, r.Dst)
		if cont == nil {
			continue
		}

		if *cont {
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

func SkipD(p *Path, next *Town) *bool {
	if len(p.Stops) == 1 {
		return PathContinue()
	}
	if p.Cost >= 30 {
		return PathDrop()
	}
	fmt.Println(p.CurrentTown().Name, next.Name)
	printPath(p)
	if p.CurrentTown().Name == "C" {
		return PathStop()
	}

	return PathContinue()
}

func PathContinue() *bool {
	t := true
	return &t
}

func PathStop() *bool {
	t := false
	return &t
}

func PathDrop() *bool {
	return nil
}
