package gsl

import (
	"fmt"
	"testing"
)

func BenchmarkBesselJ0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BesselJ0(5.0)
	}
}

func BenchmarkBesselJ1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BesselJ1(5.0)
	}
}

func BenchmarkBesselJn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BesselJn(10,5.0)
	}
}

func BenchmarkBesselJnArray(b *testing.B) {
	r := make([]float64, 10, 10)
	for i := 0; i < b.N; i++ {
		BesselJnArray(1,10,5.0,&r)
	}
}

func TestBesselJnArray(t *testing.T) {
	r := make([]float64, 10, 10)
	BesselJnArray(1,10,5.0,&r)
	for p, v := range r {
		fmt.Printf("BesselJ%d(5.0) = %v\n",p,v)
	}
}
