package libc

import "unsafe"

func Memset(ptr unsafe.Pointer, value byte, size int) {
	for i := 0; i < size; i++ {
		*(*byte)(unsafe.Pointer(uintptr(ptr) + uintptr(i))) = value
	}
}
