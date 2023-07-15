package domain

func Map[IN any, OUT any](slice []IN, f func(IN) OUT) []OUT {
	var ret []OUT
	for _, item := range slice {
		ret = append(ret, f(item))
	}
	return ret
}

func FlatMap[IN any, OUT any](slice []IN, f func(IN) []OUT) []OUT {
	var ret []OUT
	for _, item := range slice {
		ret = append(ret, f(item)...)
	}
	return ret
}

func Filter[T any](slice []T, f func(T) bool) []T {
	var ret []T
	for _, item := range slice {
		if f(item) {
			ret = append(ret, item)
		}
	}
	return ret
}
