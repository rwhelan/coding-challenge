package graph

import (
	"strconv"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

func (g *Graph) Draw(filepath string) error {
	viz := graphviz.New()
	graphViz, err := viz.Graph()
	if err != nil {
		return err
	}

	vizNodes := make(map[string]*cgraph.Node)

	for n := range g.nodes {
		gn, err := graphViz.CreateNode(n)
		if err != nil {
			return err
		}

		vizNodes[n] = gn
	}

	for _, node := range g.nodes {
		for _, edge := range node.Edges {
			e, err := graphViz.CreateEdge(
				strconv.Itoa(edge.Distance),
				vizNodes[edge.Src.Name],
				vizNodes[edge.Dst.Name],
			)
			if err != nil {
				return err
			}
			e.SetLabel(strconv.Itoa(edge.Distance))
			e.SetLen(float64(edge.Distance))
		}
	}

	if err := viz.RenderFilename(graphViz, graphviz.PNG, filepath); err != nil {
		return err
	}

	if err := graphViz.Close(); err != nil {
		return err
	}
	viz.Close()

	return nil
}
