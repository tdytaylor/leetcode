package graph

import (
	"testing"
)

func TestGraph_ShowGraph(t *testing.T) {

}

func TestGraph_insertEdge(t *testing.T) {

}

func TestGraph_insertVertex(t *testing.T) {

}

func TestInitGraph(t *testing.T) {
	graph := InitGraph(5)
	vertexs := []string{"A", "B", "C", "D", "E"}
	for _, vertex := range vertexs {
		graph.insertVertex(vertex)
	}

	graph.insertEdge(0, 1, 1)
	graph.insertEdge(0, 2, 1)
	graph.insertEdge(1, 2, 1)
	graph.insertEdge(1, 3, 1)
	graph.insertEdge(1, 4, 1)

	graph.ShowGraph()
}
