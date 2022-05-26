// Finding roots of polynomials - Newton's method

package newtons_method

import (
	"math/cmplx"
)

func dhorner(W []complex128, x complex128) (complex128, complex128) {
	n := len(W)
	y := W[0]
	z := complex(0, 0)
	for i := 1; i < n; i++ {
		z = y + z*x
		y = y*x + W[i]
	}
	return y, z
}

func PolNewton(W []complex128, p0 complex128, delta float64, max int) (complex128, float64, int) {
	var y, z, pn complex128
	var err float64
	var i int
	for i := 1; i <= max; i++ {
		y, z = dhorner(W, p0)
		pn = p0 - (y / z)
		err = cmplx.Abs(pn - p0)
		p0 = pn
		if err < delta {
			return pn, err, i
		}
	}
	return pn, err, i - 1
}
