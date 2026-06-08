package graphs

import (
	"fmt"
)

func ExampleGraph_NodesDFS() {
	g := NewGraph([][2]string{
		{"A", "C"},
		{"C", "F"},
		{"A", "D"},
		{"D", "F"},
		{"B", "D"},
		{"D", "G"},
		{"B", "E"},
		{"E", "G"},
	})

	nodes := g.NodesDFS("A")
	fmt.Println(nodes)

	// Output:
	// [A C D F G]
}
