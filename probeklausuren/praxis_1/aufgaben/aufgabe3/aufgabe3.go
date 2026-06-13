package aufgabe3

import (
	"strings"
)

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei graph.go finden Sie eine Implementierung eines Knotens in einem
//          Graphen. Die zu implementierende Funktion ist eine Methode dieses Structs.
// MAX. PUNKTE: 10

// FirstNeighbourStartingWith soll den ersten Knoten unter den Nachbarn von n liefern,
// dessen Label mit dem Präfix prefix beginnt.
// Falls keine solchen Nachbarn existieren, soll nil zurückgegeben werden.
func (n *Node) FirstNeighbourStartingWith(prefix string) *Node {
	var result *Node
	result = nil

	for _, value := range n.neighbours {
		if strings.HasPrefix(value.Label, prefix) {
			return value
		}
	}

	return result
}
