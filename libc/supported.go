package libc

var supportedLibraries = map[string]struct{}{
	"stdlib.h": {},
}

func IsSupported(library string) bool {
	_, ok := supportedLibraries[library]
	return ok
}
