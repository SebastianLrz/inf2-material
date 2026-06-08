package graphs

import (
	"fmt"
)

func ExampleGraph_NodesBFS() {
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

	nodes := g.NodesBFS("A")
	fmt.Println(nodes)

	// Output:
	// [A C D F G]
}
