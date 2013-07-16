package gsl

import (
	"fmt"
	"testing"
)

func TestErrorConstruction(t *testing.T) {
	e := GSLError{}
	s := e.Error()
	if s != "success" {
		fmt.Printf("expected string \"success\" != %s\n",s)
		t.Fail()
	}
}
