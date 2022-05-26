// Solving nonlinear equations - bisection method

package bisection_method

import (
	"fmt"
	"math"
)

func Feval(function string, parametr float64) float64 {
	switch function {
	case "sin":
		return math.Sin(parametr)
	case "cos":
		return math.Cos(parametr)
	case "tan":
		return math.Tan(parametr)
	case "exp":
		return math.Exp(parametr)
	case "log":
		return math.Log(parametr)
	case "sqrt":
		return math.Sqrt(parametr)
	default:
		return -1
	}
}

func Bisect(function string, a float64, b float64, delta float64, epsilon float64, max int64) (float64, float64, float64, int64) {
	var c, yc, err float64
	var k int64
	ya := Feval(function, a)
	yb := Feval(function, b)
	yaSign := math.Signbit(ya)
	ybSign := math.Signbit(yb)

	if yaSign == ybSign {
		fmt.Printf("Error: sign(%s(%f))==sign(%s(%f))", function, ya, function, yb)
		return -1, -1, -1, -1
	}
	for k = 1; k <= max; k++ {
		c = a + (b-a)/2
		yc = Feval(function, c)
		err = math.Abs((b - a) / 2)
		if (err < delta) || (math.Abs(yc) < epsilon) {
			return c, err, yc, k
		}
		ycSign := math.Signbit(yc)

		if ybSign == ycSign {
			b = c
			yb = yc
		} else {
			a = c
		}
	}
	return c, err, yc, k - 1
}
