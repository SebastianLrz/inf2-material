package aufgabe5

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei bintree.go finden Sie eine Implementierung eines binären Suchbaumes.
//          Die zu implementierende Funktion ist eine Methode dieses Structs.
// MAX. PUNKTE: 10

//Liste aller existierenden Pfade im Baum.
//Pfad ist String "llr" oder "rlr"
// l für "links" und r für "rechts"
func (n *Node) PathStrings() []string {
	result := []string{}
	if n.IsEmpty() {
		return result
	}
	if n.IsLeaf() {
		return []string{""}
	}
	for _, path := range n.Left.PathStrings() {
		result = append(result, "l"+path)
	}
	for _, path := range n.Right.PathStrings() {
		result = append(result, "r"+path)
	}
	return result
}
