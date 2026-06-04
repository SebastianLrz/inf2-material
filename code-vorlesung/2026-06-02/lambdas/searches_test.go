package lambdas

import "fmt"

func ExampleFirstEven() {
	l1 := []int{3, 6, 7, 9, 2, 42, 25, 33, 38}

	fmt.Println(FirstEven(l1))

	// Output:
	// 6
}

func ExampleFirstOdd() {
	l1 := []int{3, 6, 7, 9, 2, 42, 25, 33, 38}

	fmt.Println(FirstOdd(l1))

	// Output:
	// 3
}

func ExampleFirstDiv11() {
	l1 := []int{3, 6, 7, 9, 2, 42, 25, 33, 38}

	fmt.Println(FirstDiv11(l1))

	// Output:
	// 33
}

func First(l []int, pred func(int) bool) int {
	for _, el := range l {
		if pred(el) {
			return el
		}
	}
	return -1
}

func FirstEven(l []int) int {
	return First(l, func(x int) bool {
		return x%2 == 0
	})
}

func FirstOdd(l []int) int {
	p := func(x int) bool {
		return x%2 == 1
	}

	return First(l, p)
}

func ThirdOdd(l []int) int {
	counter := 0
	p := func(x int) bool {
		if x%2 == 1 {
			counter++
			if counter == 3 {
				return true
			}
		}
		return false
	}

	return First(l, p)
}

func Div11(x int) bool {
	return x%11 == 0
}

func FirstDiv11(l []int) int {
	return First(l, Div11)
}

//	func FirstEven(l []int) int {
//		for _, el := range l {
//			if el%2 == 0 {
//				return el
//			}
//		}
//		return -1
//	}

// func FirstOdd(l []int) int {
// 	for _, el := range l {
// 		if el%2 == 1 {
// 			return el
// 		}
// 	}
// 	return -1
// }

// func FirstDiv11(l []int) int {
// 	for _, el := range l {
// 		if el%11 == 0 {
// 			return el
// 		}
// 	}
// 	return -1
// }
