package addition

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes in two numbers, either integers or floats, and returns their sum.
//
// In case you don't understand how addition works, read:
// [Addition]
//
// [Addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
