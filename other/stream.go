package main

type Stream[T any] struct {
	items []T
}

func NewStream[T any](items []T) *Stream[T] {
	return &Stream[T]{items}
}
func (s *Stream[T]) Filter(predicate func(T) bool) *Stream[T] {
	items := filter(s.items, predicate)
	return &Stream[T]{items}
}

// func (s *Stream[T]) Map[U any](transform func(T) U) *Stream[U] {
// 	items := stream_map(s.items, transform)
// 	return &Stream[U]{items}
// }

// func main() {
// 	e := NewStream([]int{1, 2, 3, 4}).
// 		Filter(func(s int) bool { return s > 2 }).
// 		Filter(func(s int) bool { return s > 3 })
// 	fmt.Println(e)
// }

func filter[T any](items []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, s := range items {
		if predicate(s) {
			result = append(result, s)
		}
	}
	return result
}

func stream_map[T, U any](items []T, transform func(T) U) []U {
	result := make([]U, len(items))
	for i, s := range items {
		result[i] = transform(s)
	}
	return result
}
