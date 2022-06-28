package libc

func Ternary[T any](cond bool, true_ func() T, false_ func() T) T {
	if cond {
		return true_()
	}

	return false_()
}
