package sum

import "golang.org/x/exp/constraints"

// Add Returns the sum of two numbers
func Add[T constraints.Ordered](a, b T) T {
	return a + b
}
