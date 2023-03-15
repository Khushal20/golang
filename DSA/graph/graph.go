package graph

import "fmt"

type AdjacencyGraph struct {
	graph []*Vertix
	isDirected bool
}

type Vertix struct {
	val      int
	adjacent []*Vertix
}

func (graph *AdjacencyGraph) AddVertix(val int) {
	vertix := graph.CheckVertixExist(graph.graph, val)
	if vertix != nil {
		fmt.Printf("Vertix %v already exist.\n", val)
		return
	}
	graph.graph = append(graph.graph, &Vertix{val: val, adjacent: []*Vertix{}})
}

func (graph *AdjacencyGraph) AddEdge(from, to int) {
	fromVertix := graph.CheckVertixExist(graph.graph, from)
	toVertix := graph.CheckVertixExist(graph.graph, to)
	if fromVertix == nil && toVertix == nil{
		fmt.Printf("Invalid Edge from %v to %v.\n", from, to)
		return
	}
	tmp := graph.CheckVertixExist(fromVertix.adjacent, to)
	if tmp!=nil{
		fmt.Printf("Invalid Edge from %v to %v already exist.\n", from, to)
		return
	} 
	fromVertix.adjacent = append(fromVertix.adjacent, toVertix)
	if !graph.isDirected{
		toVertix.adjacent = append(toVertix.adjacent, fromVertix)
	}
}

func (graph *AdjacencyGraph) CheckVertixExist(verties []*Vertix, val int) *Vertix {
	for _, vertix := range verties {
		if vertix.val == val {
			return vertix
		}
	}
	return nil
}

func (graph *AdjacencyGraph) Print() {
	for _, vertix := range graph.graph {
		fmt.Printf("Vertix %v:", vertix.val)
		for _, adjVertix := range vertix.adjacent {
			fmt.Printf(" %v ", adjVertix.val)
		}
		fmt.Println("")
	}
}
