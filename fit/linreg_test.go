package gsl

import (
	"fmt"
	"testing"
	"math"
)

func withinEpsilon(x, x0, eps float64) (pass bool, delta float64) {
	delta = math.Abs(x-x0)/x0
	pass = (delta <= eps)
	return
}

func TestLinearFit(t *testing.T) {
	x := make([]float64, 10, 10)
	y := make([]float64, 10, 10)
	for pos, _ := range x {
		x[pos] = (float64)(pos)
		y[pos] = 2.0*(float64)(pos) + 4.0
	}
	fit, err := FitLinear(&x,&y,1,1)

	if err != nil {
		fmt.Printf("unexpected error in FitLinear: %v",err)
		t.Fail()
	} else {
		goodslope, slopeerr := withinEpsilon(fit.Slope, 2.0, 1e-3)
		goodicept, icepterr := withinEpsilon(fit.Y0, 4.0, 1e-3)
		if goodslope == false {
			fmt.Printf("slope error too large: %v\n",slopeerr)
			t.Fail()
		}
		if goodicept == false {
			fmt.Printf("icept error too large: %v\n",icepterr)
			t.Fail()
		}
	}
}

func TestFitWLinear(t *testing.T) {
	x := make([]float64, 10, 10)
	y := make([]float64, 10, 10)
	w := make([]float64, 10, 10)
	for pos, _ := range x {
		x[pos] = (float64)(pos)
		y[pos] = 2.0*(float64)(pos) + 4.0
		w[pos] = 0.2
	}
	fit, err := FitWLinear(&x,&y,&w,1,1,1)

	if err != nil {
		fmt.Printf("unexpected error in FitWLinear: %v",err)
		t.Fail()
	} else {
		goodslope, slopeerr := withinEpsilon(fit.Slope, 2.0, 1e-3)
		goodicept, icepterr := withinEpsilon(fit.Y0, 4.0, 1e-3)
		if goodslope == false {
			fmt.Printf("slope error too large: %v\n",slopeerr)
			t.Fail()
		}
		if goodicept == false {
			fmt.Printf("icept error too large: %v\n",icepterr)
			t.Fail()
		}
	}
}

func TestValueAt(t *testing.T) {
	x := make([]float64, 10, 10)
	y := make([]float64, 10, 10)
	for pos, _ := range x {
		x[pos] = (float64)(pos)
		y[pos] = 2.0*(float64)(pos) + 4.0
	}
	fit, err := FitLinear(&x,&y,1,1)
	if err != nil {
		fmt.Printf("unexpected error in FitLinear: %v\n",err)
		t.Fail()
	} else {
		extrapolated, ext_err, err := fit.ValueAt(35)
		if err != nil {
			fmt.Printf("unexpected error in ValueAt: %v\n",err)
		}
		goodex, _ := withinEpsilon(extrapolated, 74.0, ext_err)
		if goodex == false {
			fmt.Println("extrapolation failed.")
			t.Fail()
		} else {
			fmt.Printf("ValueAt: 74.0 == %f\n", extrapolated)
		}
	}
}
