package main

import "fmt"

// Write a function that generates a DOT representation of a graph.
// you can check the output here: https://dreampuf.github.io/GraphvizOnline/
func exercise04() {
	// struct definitions and methods in exercise03.go
	g := graph{&node{"root", nil, nil}, nil}
	n := node{"node_1", nil, nil}
	g.addNode(&n)
	n1 := node{"node_2", nil, nil}
	n.addEdge(&n1, "edge_1")
	n2 := node{"node_3", nil, nil}
	n.addEdge(&n2, "edge_2")
	n3 := node{"node_4", nil, nil}
	g.addNode(&n3)

	g.printDOT()
}

func (g *graph) printDOT() {
	fmt.Println("graph {")
	g.root.printDOT()
	fmt.Println("}")
}

func (n *node) printDOT() {
	for _, e := range n.edgesOut {
		fmt.Println("\t" + n.value.(string) + " -- " + e.n1.value.(string) + " ;")
		e.n1.printDOT()
	}
}
