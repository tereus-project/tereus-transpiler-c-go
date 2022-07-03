package libc

import "unsafe"

// void *memchr(const void *str, int c, size_t n)
func Memchr(str unsafe.Pointer, c byte, n int) unsafe.Pointer {
	for i := 0; i < n; i++ {
		if *(*byte)(unsafe.Pointer(uintptr(str) + uintptr(i))) == c {
			return unsafe.Pointer(uintptr(str) + uintptr(i))
		}
	}
	return nil
}

// int memcmp(const void *str1, const void *str2, size_t n)
func Memcmp(str1 unsafe.Pointer, str2 unsafe.Pointer, n int) int {
	for i := 0; i < n; i++ {
		if *(*byte)(unsafe.Pointer(uintptr(str1) + uintptr(i))) != *(*byte)(unsafe.Pointer(uintptr(str2) + uintptr(i))) {
			return int(*(*byte)(unsafe.Pointer(uintptr(str1) + uintptr(i)))) - int(*(*byte)(unsafe.Pointer(uintptr(str2) + uintptr(i))))
		}
	}
	return 0
}

//	void *memcpy(void *dest, const void *src, size_t n)
func Memcpy(dest unsafe.Pointer, src unsafe.Pointer, n int) unsafe.Pointer {
	for i := 0; i < n; i++ {
		*(*byte)(unsafe.Pointer(uintptr(dest) + uintptr(i))) = *(*byte)(unsafe.Pointer(uintptr(src) + uintptr(i)))
	}
	return dest
}

// void *memmove(void *dest, const void *src, size_t n)
func Memmove(dest unsafe.Pointer, src unsafe.Pointer, n int) unsafe.Pointer {
	if uintptr(dest) < uintptr(src) {
		for i := 0; i < n; i++ {
			*(*byte)(unsafe.Pointer(uintptr(dest) + uintptr(i))) = *(*byte)(unsafe.Pointer(uintptr(src) + uintptr(i)))
		}
	} else {
		for i := n - 1; i >= 0; i-- {
			*(*byte)(unsafe.Pointer(uintptr(dest) + uintptr(i))) = *(*byte)(unsafe.Pointer(uintptr(src) + uintptr(i)))
		}
	}
	return dest
}

// char *strcat(char *dest, const char *src)
func Strcat(dest unsafe.Pointer, src unsafe.Pointer) unsafe.Pointer {
	dest_len := Strlen(dest)
	src_len := Strlen(src)
	for i := 0; i < src_len; i++ {
		*(*byte)(unsafe.Pointer(uintptr(dest) + uintptr(dest_len+i))) = *(*byte)(unsafe.Pointer(uintptr(src) + uintptr(i)))
	}
	return dest
}

// char *strncat(char *dest, const char *src, size_t n)
func Strncat(dest unsafe.Pointer, src unsafe.Pointer, n int) unsafe.Pointer {
	dest_len := Strlen(dest)
	src_len := Strlen(src)
	for i := 0; i < n; i++ {
		if i < src_len {
			*(*byte)(unsafe.Pointer(uintptr(dest) + uintptr(dest_len+i))) = *(*byte)(unsafe.Pointer(uintptr(src) + uintptr(i)))
		} else {
			*(*byte)(unsafe.Pointer(uintptr(dest) + uintptr(dest_len+i))) = 0
		}
	}
	return dest
}

// char *strchr(const char *str, int c)
func Strchr(str unsafe.Pointer, c byte) unsafe.Pointer {
	for i := 0; ; i++ {
		if *(*byte)(unsafe.Pointer(uintptr(str) + uintptr(i))) == c {
			return unsafe.Pointer(uintptr(str) + uintptr(i))
		}
		if *(*byte)(unsafe.Pointer(uintptr(str) + uintptr(i))) == 0 {
			return nil
		}
	}
}

// int strcmp(const char *str1, const char *str2)
func Strcmp(str1 unsafe.Pointer, str2 unsafe.Pointer) int {
	for i := 0; ; i++ {
		if *(*byte)(unsafe.Pointer(uintptr(str1) + uintptr(i))) != *(*byte)(unsafe.Pointer(uintptr(str2) + uintptr(i))) {
			return int(*(*byte)(unsafe.Pointer(uintptr(str1) + uintptr(i)))) - int(*(*byte)(unsafe.Pointer(uintptr(str2) + uintptr(i))))
		}
		if *(*byte)(unsafe.Pointer(uintptr(str1) + uintptr(i))) == 0 {
			return 0
		}
	}
}

// int strncmp(const char *str1, const char *str2, size_t n)
func Strncmp(str1 unsafe.Pointer, str2 unsafe.Pointer, n int) int {
	for i := 0; i < n; i++ {
		if *(*byte)(unsafe.Pointer(uintptr(str1) + uintptr(i))) != *(*byte)(unsafe.Pointer(uintptr(str2) + uintptr(i))) {
			return int(*(*byte)(unsafe.Pointer(uintptr(str1) + uintptr(i)))) - int(*(*byte)(unsafe.Pointer(uintptr(str2) + uintptr(i))))
		}
		if *(*byte)(unsafe.Pointer(uintptr(str1) + uintptr(i))) == 0 {
			return 0
		}
	}
	return 0
}

// int strcoll(const char *str1, const char *str2)
func Strcoll(str1 unsafe.Pointer, str2 unsafe.Pointer) int {
	return Strcmp(str1, str2)
}

// char *strcpy(char *dest, const char *src)
func Strcpy(dest unsafe.Pointer, src unsafe.Pointer) unsafe.Pointer {
	for i := 0; ; i++ {
		*(*byte)(unsafe.Pointer(uintptr(dest) + uintptr(i))) = *(*byte)(unsafe.Pointer(uintptr(src) + uintptr(i)))
		if *(*byte)(unsafe.Pointer(uintptr(src) + uintptr(i))) == 0 {
			return dest
		}
	}
}

// char *strncpy(char *dest, const char *src, size_t n)
func Strncpy(dest unsafe.Pointer, src unsafe.Pointer, n int) unsafe.Pointer {
	for i := 0; i < n; i++ {
		*(*byte)(unsafe.Pointer(uintptr(dest) + uintptr(i))) = *(*byte)(unsafe.Pointer(uintptr(src) + uintptr(i)))
		if *(*byte)(unsafe.Pointer(uintptr(src) + uintptr(i))) == 0 {
			return dest
		}
	}
	return dest
}

// size_t strcspn(const char *str1, const char *str2)
func Strcspn(str1 unsafe.Pointer, str2 unsafe.Pointer) int {
	for i := 0; ; i++ {
		if Strchr(str2, *(*byte)(unsafe.Pointer(uintptr(str1) + uintptr(i)))) != nil {
			return i
		}
		if *(*byte)(unsafe.Pointer(uintptr(str1) + uintptr(i))) == 0 {
			return i
		}
	}
}

// char *strerror(int errnum)
func Strerror(errnum int) unsafe.Pointer {
	// TODO
	return nil
}

// size_t strlen(const char *str)
func Strlen(str unsafe.Pointer) int {
	for i := 0; ; i++ {
		if *(*byte)(unsafe.Pointer(uintptr(str) + uintptr(i))) == 0 {
			return i
		}
	}
}

// char *strpbrk(const char *str1, const char *str2)
func Strpbrk(str1 unsafe.Pointer, str2 unsafe.Pointer) unsafe.Pointer {
	for i := 0; ; i++ {
		if Strchr(str2, *(*byte)(unsafe.Pointer(uintptr(str1) + uintptr(i)))) != nil {
			return unsafe.Pointer(uintptr(str1) + uintptr(i))
		}
		if *(*byte)(unsafe.Pointer(uintptr(str1) + uintptr(i))) == 0 {
			return nil
		}
	}
}

// char *strrchr(const char *str, int c)
func Strrchr(str unsafe.Pointer, c byte) unsafe.Pointer {
	for i := Strlen(str); i >= 0; i-- {
		if *(*byte)(unsafe.Pointer(uintptr(str) + uintptr(i))) == c {
			return unsafe.Pointer(uintptr(str) + uintptr(i))
		}
	}
	return nil
}

// size_t strspn(const char *str1, const char *str2)
func Strspn(str1 unsafe.Pointer, str2 unsafe.Pointer) int {
	for i := 0; ; i++ {
		if Strchr(str2, *(*byte)(unsafe.Pointer(uintptr(str1) + uintptr(i)))) == nil {
			return i
		}
		if *(*byte)(unsafe.Pointer(uintptr(str1) + uintptr(i))) == 0 {
			return i
		}
	}
}

// char *strstr(const char *haystack, const char *needle)
func Strstr(haystack unsafe.Pointer, needle unsafe.Pointer) unsafe.Pointer {
	for i := 0; ; i++ {
		if Strncmp(haystack, needle, Strlen(needle)) == 0 {
			return unsafe.Pointer(uintptr(haystack) + uintptr(i))
		}
		if *(*byte)(unsafe.Pointer(uintptr(haystack) + uintptr(i))) == 0 {
			return nil
		}
	}
}

//	char *strtok(char *str, const char *delim)
func Strtok(str unsafe.Pointer, delim unsafe.Pointer) unsafe.Pointer {
	// TODO
	return nil
}

// size_t strxfrm(char *dest, const char *src, size_t n)
func Strxfrm(dest unsafe.Pointer, src unsafe.Pointer, n int) int {
	// TODO
	return 0
}
