package graph

import "fmt"

type Graph struct {
	vertex     []string
	edges      [][]int
	edgesValue int
}

// 创建n个节点的图
func InitGraph(n int) (graph *Graph) {
	graph = &Graph{
		vertex:     make([]string, n),
		edges:      make([][]int, n),
		edgesValue: 0,
	}
	for i := range graph.edges {
		graph.edges[i] = make([]int, n)
	}
	return
}

// 插入节点
func (g *Graph) insertVertex(value string) *Graph {
	g.vertex = append(g.vertex, value)
	g.edgesValue = g.edgesValue + 1
	return g
}

// v1 表示点的下标
func (g *Graph) insertEdge(v1, v2, weight int) *Graph {
	column := g.edges[v1]
	if column != nil {
		column[v2] = weight
	} else {
		g.edges[v1] = make([]int, cap(g.edges))
		g.edges[v1][v2] = weight
	}
	return g
}

// 显示图对应的矩阵
func (g *Graph) ShowGraph() {
	for _, val1 := range g.edges {
		fmt.Println(val1)
	}
}
