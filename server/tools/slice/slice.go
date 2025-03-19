package slice

func Find[T any](data []T, fn func(item T, index int) bool) (res T) {
	for index, item := range data {
		if fn(item, index) {
			return item
		}
	}

	return res
}

func Filter[T any](data []T, fn func(item T, index int) bool) (res []T) {
	for index, item := range data {
		if fn(item, index) {
			res = append(res, item)
		}
	}

	return res
}

func Reduce[T any, A any](data []T, fn func(item T, index int) A, acc A) A {
	for index, item := range data {
		acc = fn(item, index)
	}

	return acc
}
