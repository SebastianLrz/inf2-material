package anonyme_funktionen

// Beispiel, wie man mit zwei anonymen Funktionen und
// einer Hilfsfunktion `ForEach` eine Liste manipuliert.
func Example_anonymous() {
	// Anonyme Funktion, die eine Zahl quadriert
	quadriere := func(x int) int {
		return x * x
	}
	// Anonyme Funktion, die auf eine Zahl zwei addiert
	pluszwei := func(x int) int {
		return x + 2
	}

	// Liste erstellen
	list := []int{1, 3, 5, 7, 9}

	// `ForEach` nacheinander mit den beiden
	// anonymen verwenden.
	list = ForEach(ForEach(list, quadriere), pluszwei)
}

// Hilfsfunktion, die eine gegebene Funktion auf alle Elemente
// einer Liste anwendet.
func ForEach(l []int, f func(int) int) []int {
	for i := 0; i < len(l); i++ {
		l[i] = f(l[i])
	}
	return l
}

// Beispiel, wie das gleiche ohne anonyme Funktionen aussehen
// würde.
func Example_non_anonymous() {
	list := []int{1, 3, 5, 7, 9}

	for i := 0; i < len(list); i++ {
		list[i] = list[i] * list[i]
	}
	for i := 0; i < len(list); i++ {
		list[i] = list[i] + 2
	}
}
