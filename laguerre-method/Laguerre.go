// Finding roots of polynomials - Laguerre's method

package laguerre_method

import (
	"math/cmplx"
)

func ddhorner(W []complex128, x complex128) (complex128, complex128, complex128) {
	n := len(W)
	y := W[0]
	z := complex(0, 0)
	v := complex(0, 0)

	for i := 1; i < n; i++ {
		v = z + v*x
		z = y + z*x
		y = y*x + W[i]
	}
	v = 2.0 * v
	return y, z, v
}

func Laguerre(W []complex128, p0 complex128, delta float64, max int) (complex128, float64, int) {
	var y, z, v, A, B, C, C1, T, pn complex128
	var i int
	var err float64
	n := len(W) - 1
	for i := 1; i <= max; i++ {
		y, z, v = ddhorner(W, p0)
		if y == 0.0 {
			err = 0
			pn = p0
			return pn, err, i
		}
		A = -z / y
		B = (A * A) - (v / y)

		powerOfA := cmplx.Pow(A, A)
		nMultiplyByB := complex(real(B)*float64(n), imag(B)*float64(n))
		substractAbove := nMultiplyByB - powerOfA
		nMinusOneMultiplyByAboveSubstract := complex(real(substractAbove)*float64(n-1), imag(substractAbove)*float64(n-1))
		T = cmplx.Sqrt(nMinusOneMultiplyByAboveSubstract)        // cmplx.Sqrt( (n - 1) * ( (n * B) - (A * A) ) )
		C1 = complex(real(A+T)/float64(n), imag(A+T)/float64(n)) // (A+T) / n
		C = complex(real(A-T)/float64(n), imag(A-T)/float64(n))  // (A-T) / n
		if cmplx.Abs(C1) > cmplx.Abs(C) {
			C = C1
		}
		pn = p0 + (1.0 / C)
		err = cmplx.Abs(pn - p0)
		p0 = pn
		if err < delta {
			return pn, err, i
		}
	}
	return pn, err, i - 1
}
