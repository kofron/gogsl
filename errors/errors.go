package gsl

// #cgo LDFLAGS: -lgsl
// #include "gsl/gsl_errno.h"
import "C"

type GSLError struct {
	code C.int
}

func (e *GSLError) ErrorCode() (errno int) {
	errno = (int)(e.code)
	return
}

func (e *GSLError) Error() (what string) {
	str := C.gsl_strerror(e.code)
	what = C.GoString(str)
	return
}

func NewGSLError(code int) *GSLError {
	return &GSLError{code: (C.int)(code)}
}
