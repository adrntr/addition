// Package addition provides basic arithmetic operations.

package addition

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// Add returns the sum of two Numbers, that include float and/or int.
// More information can be found at [Addition].
//
// [Addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
