// Solving nonlinear equations - Secant method

package secant_method

import (
	"math"
	bisection_method "numericalMethodsGo/bisection-method"
)

func Secant(function string, p0 float64, p1 float64, delta float64, epsilon float64, max int) (float64, float64, float64, int) {
	var k int
	var pn, ypn, err float64

	for k := 1; k <= max; k++ {
		pn = p1 - bisection_method.Feval(function, p1)*(p1-p0)/(bisection_method.Feval(function, p1)-bisection_method.Feval(function, p0))
		err = math.Abs(pn - p1)
		p0 = p1
		p1 = pn
		ypn = bisection_method.Feval(function, pn)
		if (err < delta) || (math.Abs(ypn) < epsilon) {
			return pn, err, ypn, k
		}
	}
	return pn, err, ypn, k - 1
}
