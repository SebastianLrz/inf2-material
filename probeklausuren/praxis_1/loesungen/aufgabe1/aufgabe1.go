package aufgabe1

// AUFGABENSTELLUNG: Vervollständigen Sie die vorgegebene Funktion.
// KONTEXT: In der Datei linkedlist.go finden Sie eine Implementierung einer einfach
//          verketteten Liste. Die zu implementierende Funktion ist eine Methode dieses Structs.
// MAX. PUNKTE: 10

//true wenn Länge der Liste>5 ist
func (n *Node) LengthGreater5() bool {
	length := 0
	for !n.IsEmpty() {
		n = n.Next
		length++
	}
	return length > 5
}
