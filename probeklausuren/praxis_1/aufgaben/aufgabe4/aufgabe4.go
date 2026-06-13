package aufgabe4

import "slices"

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei graph.go finden Sie eine Implementierung eines Knotens in einem
//          Graphen. Die zu implementierende Funktion ist eine Methode dieses Structs.
// MAX. PUNKTE: 10

// ReachableNodes soll eine Liste aller von n aus erreichbaren Knoten liefern.
func (n *Node) ReachableNodes() []*Node { //mit Breitensuche

	alreadyVisited := []*Node{}

	for 



}











	/*
	alreadyVisited := []*Node{}

	//Ansatz Breitensuche: Liste von noch abzuarbeitenden Knoten
	//Suche Endet wenn Liste leer ist

	toVisit := []*Node{n}
	for len(toVisit) > 0 {
		first := toVisit[0]
		toVisit = toVisit[1:]

		if !slices.Contains(alreadyVisited, first) {
			alreadyVisited = append(alreadyVisited, first)
			toVisit = append(toVisit, first.neighbours...)
		}

	}

	return alreadyVisited
}
	*/