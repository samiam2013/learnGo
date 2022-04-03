package intermediate

import "fmt"

type graph struct {
	root  *node
	nodes []node
}

type node struct {
	value    interface{}
	edgesIn  []edge
	edgesOut []edge
}

type edge struct {
	value interface{}
	n     *node
	n1    *node
}

func (g *graph) addNode(n *node) {
	g.root.addEdge(n, nil)
}

func (n *node) addEdge(n1 *node, value interface{}) {
	newEdge := edge{value, n, n1}
	n.edgesOut = append(n.edgesOut, newEdge)
	n1.edgesIn = append(n1.edgesIn, newEdge)
}

func (n *node) delete() {
	for _, e := range n.edgesIn {
		e.delete()
	}
	for _, e := range n.edgesOut {
		e.delete()
	}
	*n = node{"", nil, nil}
}

func (e *edge) delete() {
	e.n.edgesOut = removeSliceEdge(e.n.edgesOut, e)
	e.n.edgesIn = removeSliceEdge(e.n1.edgesIn, e)
}

func removeSliceEdge(edges []edge, edge *edge) []edge {
	for i := 0; i < len(edges); i++ {
		if edges[i] == *edge {
			edges[i] = edges[len(edges)-1] // Copy last element to index i.
			edges = edges[:len(edges)-1]   // Truncate slice.
			break
		}
	}
	// slice header pointer moves on truncate,
	//  so value has to be returned despite
	//  oterwise pass-by-reference behavior of slices
	return edges
}

func (n *node) print() {
	fmt.Println("node value:", n.value)
	for _, e := range n.edgesOut {
		if e.value != nil {
			fmt.Println("edge value:", e.value)
		}
		e.n1.print()
	}
}

// Implement a data structure for graphs that allows modification
//  (insertion, deletion). It should be possible to store values
//  at edges and nodes. It might be easiest to use a dictionary of
//  (node, edgelist) to do this.
func exercise03() {
	g := graph{&node{"root", nil, nil}, nil}
	n := node{"value1", nil, nil}
	g.addNode(&n)
	n1 := node{"value2", nil, nil}
	n.addEdge(&n1, "edgeVal1")
	fmt.Println("printing from node n")
	n.print()
	fmt.Println("printing after n1.delete()")
	n1.delete()
	n.print()
}
