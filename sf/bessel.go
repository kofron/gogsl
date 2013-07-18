/*
bessel.go
golang wrappers for bessel functions in gsl.
*/

package sf

// #cgo LDFLAGS: -lgsl
// #include "gsl/gsl_sf_bessel.h"
import "C"

import (
	"unsafe"
	"github.com/kofron/gogsl"
)

func BesselJ0(x float64) (result float64) {
	result = (float64)(C.gsl_sf_bessel_J0((C.double)(x)))
	return
}

func BesselJ1(x float64) (result float64) {
	result = (float64)(C.gsl_sf_bessel_J0((C.double)(x)))
	return 
}

func BesselJn(order int, x float64) (result float64) {
	result = (float64)(C.gsl_sf_bessel_Jn((C.int)(order),(C.double)(x)))
	return
}

func BesselJnArray(nmin, nmax int, x float64, result *[]float64) (e error) {
	resptr := (*C.double)((unsafe.Pointer)(&(*result)[0]))
	er := (C.gsl_sf_bessel_Jn_array((C.int)(nmin), 
		(C.int)(nmax), 
		(C.double)(x), 
		resptr))
	if er == (C.int)(0) {
		e = nil
	} else {
		e = gsl.NewGSLError(int(er))
	}

	return
}
