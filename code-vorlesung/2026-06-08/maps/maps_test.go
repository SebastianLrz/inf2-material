package maps

import (
	"fmt"
	"maps"
	"slices"
)

func Example() {
	m := make(map[string]bool)
	m["a"] = true
	m["b"] = true
	m["c"] = true

	for key, value := range m {
		fmt.Println(key, value)
	}
	fmt.Println("d", m["d"])

	// Output:
	// a true
	// b true
	// c true
	// d false
}

func Example_iterate() {
	m := make(map[string]bool)
	m["a"] = true
	m["b"] = true
	m["c"] = true

	for key := range m {
		fmt.Println(key)
		m["d"] = true
	}

	// Output:
}

func Example_iterate_2() {
	m := make(map[string]bool)
	m["a"] = true
	m["b"] = true
	m["c"] = true

	for i := 0; i < len(m); i++ {
		keys := slices.Collect(maps.Keys(m))
		slices.Sort(keys)
		fmt.Println(keys[i])
		m["d"] = true
	}

	// Output:
}
