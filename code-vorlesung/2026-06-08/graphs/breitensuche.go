package graphs

import (
	"maps"
	"slices"
)

// Breitensuche nach allen erreichbaren Knoten von start aus.
func (g *Graph) NodesBFS(start string) []string {
	visited := make(map[string]bool)
	queue := []string{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if !visited[node] {
			visited[node] = true
			queue = append(queue, g.edges[node]...)
		}
	}

	result := slices.Collect(maps.Keys(visited))
	slices.Sort(result)
	return result
}
