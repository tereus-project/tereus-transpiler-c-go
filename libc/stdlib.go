package libc

import (
	"sync"
	"unsafe"
)

var (
	memoryAddresses map[uint64]interface{}
	memoryMutex     sync.Mutex
)

func init() {
	memoryAddresses = make(map[uint64]interface{})
}

func Malloc(size int) unsafe.Pointer {
	memory := make([]byte, size)
	address := uint64(uintptr(unsafe.Pointer(&memory[0])))

	memoryMutex.Lock()
	defer memoryMutex.Unlock()

	memoryAddresses[address] = memory
	return unsafe.Pointer(&memory[0])
}

func Free(pointer unsafe.Pointer) {
	address := uint64(uintptr(pointer))

	memoryMutex.Lock()
	defer memoryMutex.Unlock()

	delete(memoryAddresses, address)
}
