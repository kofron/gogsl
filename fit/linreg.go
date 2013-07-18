package fit

// #cgo LDFLAGS: -lgsl
// #include "gsl/gsl_fit.h"
import "C"

import (
	"fmt"
	"unsafe"
	"github.com/kofron/gogsl/errors"
)

/* A simple storage container for Linear Fit results.  Note that not all
 * fields are necessarily populated by a given fitting routine - check 
 * the documentation to be sure.
 */
type LinearFit struct {
	Y0, Slope, Cov00, Cov01, Cov11, SumSq, ChiSq float64
}

/* ValueAt wraps gsl_fit_linear_est from gsl_fit.h.
 * A LinearFit can be interrogated for the value of a result at a particular
 * point x.  All of the parameters mentioned below are taken from the 
 * LinearFit.
 * From the GSL manual:
 * This function uses the best-fit linear regression coefficients Y0, Slope 
 * and their covariance Cov00, Cov01, Cov11 to compute the fitted function 
 * y and its standard deviation y_err for themodel 
 * Y = Y0 + Slope* X at the point x.
 */
func (f *LinearFit) ValueAt(x float64) (y, y_err float64, e error) {
	er := C.gsl_fit_linear_est((C.double)(x),
		(C.double)(f.Y0),
		(C.double)(f.Slope),
		(C.double)(f.Cov00),
		(C.double)(f.Cov01),
		(C.double)(f.Cov11),
		(*C.double)(&y),
		(*C.double)(&y_err))
	if er != (C.int)(0) {
		e = gsl.NewGSLError((int)(er))
	}
	return
}

/* FitLinear wraps gsl_fit_linear from gsl_fit.h.
 * All of the parameters mentioned below are returned from the function 
 * in a *LinearFit.
 * From the GSL manual:
 * This function computes the best-fit linear regression coefficients 
 * (Y0,Slope) of the model Y = Y0 + Slope*X for the dataset (x, y), two
 * vectors 
 * of identical length n with strides xstride and ystride. The errors 
 * on y are assumed unknown so the variance-covariance matrix for the 
 * parameters (Y0, Slope) is estimated from the scatter of the points 
 * around the best-fit line and returned via the parameters 
 * (Cov00, Cov01, Cov11). 
 * The sum of squares of the residuals from the best-fit line is returned 
 * in Sumsq.
 */
func FitLinear(x,y *[]float64, stridex, stridey uint) (f *LinearFit, e error){
	f = &LinearFit{}
	xptr := (*C.double)(unsafe.Pointer(&(*x)[0]))
	yptr := (*C.double)(unsafe.Pointer(&(*y)[0]))
	er := C.gsl_fit_linear(
		xptr,
		(C.size_t)(stridex),
		yptr,
		(C.size_t)(stridey),
		(C.size_t)(len(*x)),
		(*C.double)(&f.Y0),
		(*C.double)(&f.Slope),
		(*C.double)(&f.Cov00),
		(*C.double)(&f.Cov01),
		(*C.double)(&f.Cov11),
		(*C.double)(&f.SumSq))
	if er == (C.int)(0) {
		e = nil
		fmt.Printf("Sum of Squares: %v\n",f.SumSq)
	} else {
		e = gsl.NewGSLError(int(er))
	}
	return
}

/* FitWLinear wraps gsl_fit_wlinear.h 
 * Paraphrasing the GSL manual: 
 * This function computes the best-fit linear regression coefficients 
 * (Y0,Slope) of the model Y = Y0 + Slope*X for the weighted dataset 
 * (x, y), two vectors of length n with strides xstride and ystride. 
 * The vector w, of length n and stride wstride, specifies the weight of 
 * each datapoint. The weight is the reciprocal of the variance for each 
 * datapoint in y.  The covariance matrix for the parameters (Y0, Slope) is 
 * computed using the weights and returned via the parameters 
 * (Cov00, Cov01, Cov11). The weighted sum of squares of the residuals from 
 * the best-fit line, χ2, is returned in ChiSq.
 */
func FitWLinear(x,y,w *[]float64, sx, sy, sw uint) (f *LinearFit, e error) {
		f = &LinearFit{}
	xptr := (*C.double)(unsafe.Pointer(&(*x)[0]))
	yptr := (*C.double)(unsafe.Pointer(&(*y)[0]))
	wptr := (*C.double)(unsafe.Pointer(&(*w)[0]))
	er := C.gsl_fit_wlinear(
		xptr,
		(C.size_t)(sx),
		wptr,
		(C.size_t)(sw),
		yptr,
		(C.size_t)(sy),
		(C.size_t)(len(*x)),
		(*C.double)(&f.Y0),
		(*C.double)(&f.Slope),
		(*C.double)(&f.Cov00),
		(*C.double)(&f.Cov01),
		(*C.double)(&f.Cov11),
		(*C.double)(&f.ChiSq))
	if er == (C.int)(0) {
		e = nil
		fmt.Printf("Chi Squared: %v\n",f.ChiSq)
	} else {
		e = gsl.NewGSLError(int(er))
	}
	return
}
