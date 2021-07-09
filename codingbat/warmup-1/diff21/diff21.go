package diff21

import "math"

func Diff21(n int) (result int) {
	result = int(math.Abs(float64(n - 21)))
	if n > 21 {
		result *= 2
	}
	return
}
