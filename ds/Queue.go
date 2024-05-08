package ds

type Queue[T any] struct {
	Data []T
}

func (q *Queue[T]) Length() int {
	return len(q.Data)
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		Data: []T{},
	}
}

func (q *Queue[T]) Push(val T) {
	q.Data = append(q.Data, val)
}

func (q *Queue[T]) Pop() T {

	var val T
	val, q.Data = q.Data[0], q.Data[1:]
	return val
}
