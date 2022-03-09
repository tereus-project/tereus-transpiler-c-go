package libc

var supportedLibraries = map[string]struct{}{
	"stdlib": {},
}

func IsSupported(library string) bool {
	_, ok := supportedLibraries[library]
	return ok
}
