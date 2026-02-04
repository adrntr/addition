// Package addition provides basic arithmetic operations.

package addition

// Add returns the sum of two integers.
// If the result overflows the int type, it will wrap around according to Go's spec.
// More information can be found at [Addition].
//
// [Addition]: https://www.mathsisfun.com/numbers/addition.html
func Add(a, b int) int {
	return a + b
}
