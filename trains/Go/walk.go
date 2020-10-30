package main

type contFunc func(p *Path, next *Node) *bool

func walkr(p *Path, f contFunc) *PathList {
	all := NewPathList()

	for _, e := range p.CurrentNode().Edges {
		cont := f(p, e.Dst)
		if cont == nil {
			continue
		}

		if *cont {
			np := &Path{
				Nodes: append(p.Nodes, e.Dst),
				Edges: append(p.Edges, e),
				Cost:  p.Cost + e.Distance,
			}
			all.Add(walkr(np, f))

		} else {
			// Dedup
			if all.Len() == 0 ||
				all.Len() >= 1 &&
					!(all.Last() == p) {
				all.Append(p)
			}
		}
	}

	return all
}

func walk(all *PathList, p *Path, f contFunc) {
	for _, e := range p.CurrentNode().Edges {
		cont := f(p, e.Dst)
		if cont == nil {
			continue
		}

		if *cont {
			np := &Path{
				Nodes: append(p.Nodes, e.Dst),
				Edges: append(p.Edges, e),
				Cost:  p.Cost + e.Distance,
			}
			walk(all, np, f)

		} else {
			// Dedup
			if all.Len() == 0 ||
				all.Len() >= 1 &&
					!(all.Last() == p) {
				all.Append(p)
			}
		}
	}
}

func SkipD(p *Path, next *Node) *bool {
	if len(p.Nodes) == 1 {
		return PathContinue()
	}
	if p.Cost >= 30 {
		return PathDrop()
	}
	//fmt.Println(p.CurrentNode().Name, next.Name)
	//PrintPath(p)
	if p.CurrentNode().Name == "C" {
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
