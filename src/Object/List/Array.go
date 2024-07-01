package List

type Array[T any] struct {
	slice []T
}

func (a Array[T]) Map(fn func(value T, index int) T) Array[T] {
	for i, v := range a.slice {
		a.slice[i] = fn(v, i)
	}
	return a
}

func (a Array[T]) Filter(fn func(value T) bool) Array[T] {
	var local = []T{}
	for _, v := range a.slice {
		if fn(v) {
			local = append(local, v)
		}
	}
	a.slice = local
	return a
}

func (a Array[T]) Reduce() Array[T] {
	// TODO: Implement Reduce
	return a
}

func (a Array[T]) Build() []T {
	return a.slice
}

func NewArray[T any](slice []T) Array[T] {
	return Array[T]{slice: slice}
}
