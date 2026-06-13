package aufgabe4

import "slices"

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei graph.go finden Sie eine Implementierung eines Knotens in einem
//          Graphen. Die zu implementierende Funktion ist eine Methode dieses Structs.
// MAX. PUNKTE: 10

//Liste aller Knoten liefern,diein genau d Shrittn
// von n aus erreichbar sind.
//keine Dplikte enthalten,muss nicht sortiert sein
//n selbst ist in 0 Schritten von n aus erreichba
//seine direkten Nchbrn sind in 1 Schrtt errchbar
func (n *Node) NodesWithDistance(d int) []*Node {
	result := []*Node{n}
	for ; d > 0; d-- {
		frontier := []*Node{}
		for _, n := range result {
			for _, n := range n.neighbours {
				if !slices.Contains(frontier, n) {
					frontier = append(frontier, n)
				}
			}
		}
		result = frontier
	}
	return result
}
