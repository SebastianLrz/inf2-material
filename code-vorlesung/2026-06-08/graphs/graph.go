package graphs

type Graph struct {
	edges map[string][]string
}

func NewGraph(edges [][2]string) *Graph {
	m := make(map[string][]string)
	for _, edge := range edges {
		m[edge[0]] = append(m[edge[0]], edge[1])
	}
	return &Graph{edges: m}
}
