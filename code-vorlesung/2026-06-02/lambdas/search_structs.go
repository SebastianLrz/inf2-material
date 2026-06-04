package lambdas

import "fmt"

type Searcher struct {
	pred func(int) bool
}

func NewSearcher(p func(int) bool) *Searcher {
	return &Searcher{pred: p}
}

func (s Searcher) Find(l []int) int {
	for _, el := range l {
		if s.pred(el) {
			return el
		}
	}
	return -1
}

func ExampleSearcher() {
	s := NewSearcher(func(x int) bool {
		return x%5 == 0
	})

	fmt.Println(s.Find([]int{2, 4, 6, 8, 9, 12, 15, 32}))
	fmt.Println(s.Find([]int{1, 3, 7, 9}))
	fmt.Println(s.Find([]int{45, 35, 25, 15}))

	// Output:
	// 15
	// -1
	// 45
}
