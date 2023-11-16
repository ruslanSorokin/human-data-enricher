package util

func MustXXX[T, U any](a T, b U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return a, b
}

func MustXX[T any](a T, err error) T {
	if err != nil {
		panic(err)
	}
	return a
}

func MustX(err error) {
	if err != nil {
		panic(err)
	}
}
