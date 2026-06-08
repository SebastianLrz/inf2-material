package graphs

import "slices"

// Tiefensuche nach allen erreichbaren Knoten von start aus.
func (g *Graph) NodesDFS(start string) []string {
	visited := []string{start}
	for _, n := range g.edges[start] {
		visited = append(visited, g.NodesDFS(n)...)
	}
	slices.Sort(visited)
	return slices.Compact(visited)
}
