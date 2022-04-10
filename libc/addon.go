package libc

import (
	"unsafe"

	"golang.org/x/exp/constraints"
)

func PreInc[T constraints.Integer](v *T) T {
	*v++
	return *v
}

func PreIncPtr[T any](v **T) *T {
	*v = (*T)(unsafe.Add(unsafe.Pointer(*v), 1))
	return *v
}

func PostInc[T constraints.Integer](v *T) T {
	old := *v
	*v++
	return old
}

func PostIncPtr[T any](v **T) *T {
	old := *v
	*v = (*T)(unsafe.Add(unsafe.Pointer(*v), 1))
	return old
}

func PreDec[T constraints.Integer](v *T) T {
	*v++
	return *v
}

func PreDecPtr[T any](v **T) *T {
	*v = (*T)(unsafe.Add(unsafe.Pointer(*v), -1))
	return *v
}

func PostDec[T constraints.Integer](v *T) T {
	old := *v
	*v++
	return old
}

func PostDecPtr[T any](v **T) *T {
	old := *v
	*v = (*T)(unsafe.Add(unsafe.Pointer(*v), -1))
	return old
}
